package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/chaossat/tiktak/common"
	"github.com/chaossat/tiktak/db"
	"github.com/chaossat/tiktak/middleware"
	"github.com/chaossat/tiktak/model"
	"github.com/chaossat/tiktak/service/comment/pb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	log.Println("正在启动Comment服务......")
	InitConfig()
	common.InitDB()
	//初始化grpc实例
	grpcServer := grpc.NewServer()

	//注册服务
	pb.RegisterCommentActionServer(grpcServer, new(CommentActionHandler))
	pb.RegisterCommentListServer(grpcServer, new(CommentListHandler))

	//设置监听
	listen, err := net.Listen("tcp", ":10005")
	if err != nil {
		log.Println("注册服务启动监听失败")
	}
	defer listen.Close()

	//启动服务
	grpcServer.Serve(listen)
}

func InitConfig() {
	workDir, _ := os.Getwd() //获取当前工作路径，非文件路径，以终端显示路径为准
	viper.SetConfigName("config.yml")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	fmt.Println("configPath:", workDir+"/config")
	err := viper.ReadInConfig()
	fmt.Println(
		"host:", viper.GetString("datasource.mysql.host"), "\r\n",
		"port:", viper.GetString("datasource.mysql.port"), "\r\n",
		"database:", viper.GetString("datasource.mysql.database"), "\r\n",
		"username:", viper.GetString("datasource.mysql.username"), "\r\n",
		"passowrd:", viper.GetString("datasource.mysql.password"), "\r\n",
		"charset:", viper.GetString("datasource.mysql.charset"),
	)
	if err != nil {
		panic(err)
	}
}

//GetPbUser根据用户id返回符合protobuf文件的user
func GetPbUser(uid, myid int) *pb.User {
	user, err := db.UserInfoById(uid)
	if err != nil {
		fmt.Println("获取用户信息时失败", "userid:", uid, err.Error())
		return nil
	}
	followCount, err := db.FollowCountByID(int(user.ID))
	if err != nil {
		fmt.Println("获取用户关注数时失败:" + err.Error())
		return nil
	}
	followerCount, err := db.FollowerCountByID(int(user.ID))
	if err != nil {
		fmt.Println("获取用户粉丝数时失败:" + err.Error())
		return nil
	}
	isFollow := false
	if myid != uid {
		isFollow = db.IsFollow(model.User{ID: int64(myid)}, model.User{ID: int64(uid)})
	}
	pbuser := pb.User{
		Id:            &user.ID,
		Name:          &user.Username,
		FollowCount:   &followCount,
		FollowerCount: &followerCount,
		IsFollow:      &isFollow,
	}
	return &pbuser
}

type CommentActionHandler struct {
	pb.UnimplementedCommentActionServer
}

//CommentAction:grpc的CommentAction接口
func (co *CommentActionHandler) CommentAction(ctx context.Context, req *pb.DouyinCommentActionRequest) (*pb.DouyinCommentAcitonResponse, error) {
	// token验证
	payload, err := middleware.CheckToken(req.GetToken())
	if err != nil {
		return CommentActionResponse(-1, "解析token错误:"+err.Error(), 0, ErrorUser(), "", 0), nil
	}
	// if payload.UserID != *req.UserId {
	// 	fmt.Println("用户id:", *req.UserId, "token id:", payload.UserID)
	// 	return CommentActionResponse(-2, "用户id与token不匹配", 0, ErrorUser(), "", 0), nil
	// }
	//查询视频是否存在
	b, err := db.IsVideoExist(int(req.GetVideoId()))
	if err != nil {
		return CommentActionResponse(-3, "查询视频时发生错误:"+err.Error(), 0, ErrorUser(), "", 0), nil
	}
	if !b {
		return CommentActionResponse(-4, "视频不存在!", 0, ErrorUser(), "", 0), nil
	}
	//判断操作类型
	if req.GetActionType() == 2 {
		//删除操作
		err := db.DeleteComment(req.GetCommentId())
		if err != nil {
			return CommentActionResponse(-5, "删除评论时发生错误:"+err.Error(), 0, ErrorUser(), "", 0), nil
		}
		return CommentActionResponse(0, "评论删除成功", 0, ErrorUser(), "", 0), nil
	} else if req.GetActionType() == 1 {
		//创建操作
		pbuser := GetPbUser(int(payload.UserID), int(payload.UserID))
		if pbuser == nil {
			return CommentActionResponse(-6, "获取用户信息时失败", 0, ErrorUser(), "", 0), nil
		}
		comment := model.Comment{
			Content:    req.GetCommentText(),
			UserID:     payload.UserID,
			VideoID:    req.GetVideoId(),
			UpdateTime: time.Now().Unix(),
		}
		err = db.CreateComment(&comment)
		if err != nil {
			return CommentActionResponse(-7, "创建评论时失败:"+err.Error(), 0, ErrorUser(), "", 0), nil
		}
		return CommentActionResponse(0, "发表评论成功", int(comment.ID), pbuser, comment.Content, comment.UpdateTime), nil
	} else {
		return CommentActionResponse(1, "非法的请求类型", 0, ErrorUser(), "", 0), nil
	}
}

//创建一个空user并返回，用默认值填充了必填字段
func ErrorUser() *pb.User {
	id := int64(0)
	name := ""
	isfollow := false
	errorUser := pb.User{
		Id:       &id,
		Name:     &name,
		IsFollow: &isfollow,
	}
	return &errorUser
}

//CommentActionResponse:返回评论处理信息
func CommentActionResponse(code int, msg string, commentID int, user *pb.User, content string, cdate int64) *pb.DouyinCommentAcitonResponse {
	codeResponse := int32(code)
	cid := int64(commentID)
	date := time.Unix(cdate, 0).String()[5:10]
	return &pb.DouyinCommentAcitonResponse{
		StatusCode: &codeResponse,
		StatusMsg:  &msg,
		Comment: &pb.Comment{
			Id:         &cid,
			User:       user,
			Content:    &content,
			CreateDate: &date,
		},
	}
}

type CommentListHandler struct {
	pb.UnimplementedCommentListServer
}

//CommentList:grpc的CommentList接口
func (co *CommentListHandler) CommentList(ctx context.Context, req *pb.DouyinCommentListRequest) (*pb.DouyinCommentListResponse, error) {
	// token验证
	payload, err := middleware.CheckToken(req.GetToken())
	if err != nil {
		return CommentListResponse(-1, "校验token失败", nil), nil
	}
	//查询视频是否存在
	b, err := db.IsVideoExist(int(req.GetVideoId()))
	if err != nil {
		return CommentListResponse(-2, "查询视频时发生错误:"+err.Error(), nil), nil
	}
	if !b {
		return CommentListResponse(-3, "视频不存在!", nil), nil
	}
	//获取评论
	comments, err := db.CommentsByVID(int(req.GetVideoId()))
	if err != nil {
		return CommentListResponse(-4, "获取评论列表失败", nil), nil
	}
	pbcomments := make([]*pb.Comment, len(comments))
	for i, j := range comments {
		date := time.Unix(j.UpdateTime, 0).String()[5:10]
		user := GetPbUser(int(j.UserID), int(payload.UserID))
		id := j.ID
		content := j.Content
		if user == nil {
			return CommentListResponse(-5, "获取评论用户信息失败", nil), nil
		}
		pbcomments[i] = &pb.Comment{
			Id:         &id,
			User:       user,
			Content:    &content,
			CreateDate: &date,
		}
	}
	return CommentListResponse(0, "获取评论列表成功", pbcomments), nil
}

//CommentListResponse:返回评论列表处理信息
func CommentListResponse(code int, msg string, list []*pb.Comment) *pb.DouyinCommentListResponse {
	scode := int32(code)
	return &pb.DouyinCommentListResponse{
		StatusCode:  &scode,
		StatusMsg:   &msg,
		CommentList: list,
	}
}
