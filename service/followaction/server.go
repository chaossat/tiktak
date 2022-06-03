package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/chaossat/tiktak/common"
	"github.com/chaossat/tiktak/db"
	"github.com/chaossat/tiktak/middleware"
	"github.com/chaossat/tiktak/service/followaction/pb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type FollowAction struct {
	pb.UnimplementedFollowActionServer
}

func (f *FollowAction) FollowAction(ctx context.Context, req *pb.DouyinRelationActionRequest) (*pb.DouyinRelationActionResponse, error) {
	//定义返回值
	var statuscode int32
	var statusmsg string
	if *req.ToUserId == *req.UserId {
		statuscode = -1
		statusmsg = "不能关注自己"
		return &pb.DouyinRelationActionResponse{
			StatusCode: &statuscode,
			StatusMsg:  &statusmsg,
		}, nil
	}
	//验证token , 错误代码1
	token := *req.Token
	_, err := middleware.CheckToken(token)
	if err != nil {
		if err != nil {
			log.Println("解析用户token失败", err.Error())
			statuscode = 1
			statusmsg = "解析用户token失败"
			return &pb.DouyinRelationActionResponse{
				StatusCode: &statuscode,
				StatusMsg:  &statusmsg,
			}, err
		}
	}
	//验证user_id , 错误代码2
	userinf, err := db.UserInfoById(int(*req.UserId))
	log.Println(*req.UserId)
	if err != nil || userinf.ID == 0 {
		statuscode = 2
		statusmsg = "验证user_id失败"
		return &pb.DouyinRelationActionResponse{
			StatusCode: &statuscode,
			StatusMsg:  &statusmsg,
		}, err
	}
	//验证to_user_id , 错误代码3
	touserinf, err := db.UserInfoById(int(*req.ToUserId))
	if err != nil || touserinf.ID == 0 {
		statuscode = 3
		statusmsg = "验证to_user_id失败"
		return &pb.DouyinRelationActionResponse{
			StatusCode: &statuscode,
			StatusMsg:  &statusmsg,
		}, err
	}
	//操作action_type , 错误代码4
	log.Println("进行操作")
	if *req.ActionType == int32(1) {
		//关注
		err = db.Follow(userinf, touserinf)
	} else if *req.ActionType == int32(2) {
		//取消关注
		err = db.DelFollow(userinf, touserinf)
	}
	if err != nil {
		statuscode = 4
		statusmsg = "验证user_id失败"
		return &pb.DouyinRelationActionResponse{
			StatusCode: &statuscode,
			StatusMsg:  &statusmsg,
		}, err
	}

	//成功
	statuscode = 0
	statusmsg = "关注成功"
	return &pb.DouyinRelationActionResponse{
		StatusCode: &statuscode,
		StatusMsg:  &statusmsg,
	}, err
}

func main() {
	log.Println("正在启动FollowAction服务......")
	InitConfig()
	common.InitDB()
	common.InitRedis()
	//初始化grpc实例
	grpcServer := grpc.NewServer()

	//注册服务
	pb.RegisterFollowActionServer(grpcServer, new(FollowAction))

	//设置监听
	listen, err := net.Listen("tcp", ":12358")
	if err != nil {
		log.Println("注册服务启动监听失败")
	}
	defer listen.Close()

	//启动服务
	grpcServer.Serve(listen)
}

func InitConfig() {
	workDir, _ := os.Getwd() //获取当前工作路径，非文件路径，以终端显示路径为准
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
