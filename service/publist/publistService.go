package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/CodyGuo/godaemon"
	"github.com/chaossat/tiktak/common"
	"github.com/chaossat/tiktak/db"
	"github.com/chaossat/tiktak/middleware"
	"github.com/chaossat/tiktak/oss"
	feedmodel "github.com/chaossat/tiktak/service/feed/model"
	"github.com/chaossat/tiktak/service/publist/pb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type PublistServer struct {
	pb.UnimplementedPublishServer
}

// 获取发布的video列表
func (this *PublistServer) PublishVideo(ctx context.Context, req *pb.DouyinPublishListRequest) (*pb.DouyinPublishListResponse, error) {
	// token验证
	_, err := middleware.CheckToken(*req.Token)
	if err != nil {
		var code int32 = -1
		var msg string = "验证失败！"
		return &pb.DouyinPublishListResponse{
			StatusCode: &code,
			StatusMsg:  &msg,
			VideoList:  nil,
		}, nil
	}
	userinf, err := db.UserInfoById(int(*req.UserId))
	if err != nil {
		var code int32 = -2
		var msg string = "查询用户信息失败！"
		return &pb.DouyinPublishListResponse{
			StatusCode: &code,
			StatusMsg:  &msg,
			VideoList:  nil,
		}, nil
	}
	if userinf.ID == 0 {
		var code int32 = -3
		var msg string = "对应id的用户不存在！"
		return &pb.DouyinPublishListResponse{
			StatusCode: &code,
			StatusMsg:  &msg,
			VideoList:  nil,
		}, nil
	}
	videos, err := db.VideoedByID(int(*req.UserId))
	if err != nil {
		var code int32 = -4
		var msg string = "查询用户视频失败！"
		return &pb.DouyinPublishListResponse{
			StatusCode: &code,
			StatusMsg:  &msg,
			VideoList:  nil,
		}, nil
	}
	isfollow := false
	followCount, err := db.FollowCountByID(int(userinf.ID))
	if err != nil {
		var code int32 = -5
		var msg string = "查询关注数量失败！"
		return &pb.DouyinPublishListResponse{
			StatusCode: &code,
			StatusMsg:  &msg,
			VideoList:  nil,
		}, nil
	}
	followerCount, err := db.FollowerCountByID(int(userinf.ID))
	if err != nil {
		var code int32 = -6
		var msg string = "查询粉丝数量失败！"
		return &pb.DouyinPublishListResponse{
			StatusCode: &code,
			StatusMsg:  &msg,
			VideoList:  nil,
		}, nil
	}
	var user = pb.User{
		Id:            &userinf.ID,
		Name:          &userinf.Username,
		FollowCount:   &followCount,
		FollowerCount: &followerCount,
		IsFollow:      &isfollow,
	}
	var vide_ans = make([]*pb.Video, len(videos))
	for i := 0; i < len(videos); i++ {
		var isfabvorite = false
		playURL := oss.GetURL(videos[i].Location)
		coverURL := oss.GetURL(videos[i].Cover_location)
		favoritecnt, err := feedmodel.GetFavoriteCount(videos[i].ID)
		if err != nil {
			log.Println("获取视频的点赞个数错误", err.Error())
			var code int32 = -7
			var msg string = "获取视频的点赞个数错误！"
			return &pb.DouyinPublishListResponse{
				StatusCode: &code,
				StatusMsg:  &msg,
				VideoList:  nil,
			}, nil
		}
		cCount := int64(0)
		vide_ans[i] = &pb.Video{
			Id:            &videos[i].ID,
			Author:        &user,
			PlayUrl:       &playURL,
			CoverUrl:      &coverURL,
			FavoriteCount: &favoritecnt,
			CommentCount:  &cCount,
			IsFavorite:    &isfabvorite,
			Title:         &videos[i].Title,
		}
	}
	if err != nil {
		var code int32 = -8
		var msg string = "视频查找失败！"
		return &pb.DouyinPublishListResponse{
			StatusCode: &code,
			StatusMsg:  &msg,
			VideoList:  nil,
		}, nil
	}
	var code int32 = 0
	var msg string = "验证成功！"
	res := &pb.DouyinPublishListResponse{
		StatusCode: &code,
		StatusMsg:  &msg,
		VideoList:  vide_ans,
	}
	return res, nil
}

func main() {
	log.Println("正在启动Publist服务......")
	InitConfig()
	common.InitDB()
	feedmodel.InitRedis()
	//初始化grpc实例
	grpcServer := grpc.NewServer()

	//注册服务
	pb.RegisterPublishServer(grpcServer, new(PublistServer))

	//设置监听
	listen, err := net.Listen("tcp", ":10002")
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
