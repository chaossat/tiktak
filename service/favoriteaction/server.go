package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/chaossat/tiktak/middleware"
	"github.com/chaossat/tiktak/service/favoriteaction/model"
	"github.com/chaossat/tiktak/service/favoriteaction/pb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type FavoriteAction struct {
}

func (this *FavoriteAction) GetFavoriteAction(ctx context.Context, req *pb.DouyinFavoriteActionRequest) (*pb.DouyinFavoriteActionResponse, error) {
	var statuscode int32
	var statusmsg string
	token := *req.Token
	_, err := middleware.CheckToken(token)
	if err != nil {
		if err != nil {
			log.Println("解析用户token失败", err.Error())
			statuscode = 1
			statusmsg = "解析用户token失败"
			return &pb.DouyinFavoriteActionResponse{
				StatusCode: &statuscode,
				StatusMsg:  &statusmsg,
			}, nil
		}
	}
	if *req.ActionType == 1 {
		_, err := model.GiveLike(*req.UserId, *req.VideoId)
		if err != nil {
			log.Println("点赞操作失败", err.Error())
			statuscode = 1
			statusmsg = "点赞操作失败"
			return &pb.DouyinFavoriteActionResponse{
				StatusCode: &statuscode,
				StatusMsg:  &statusmsg,
			}, nil
		}
		statuscode = 0
		statusmsg = "点赞成功"
		return &pb.DouyinFavoriteActionResponse{
			StatusCode: &statuscode,
			StatusMsg:  &statusmsg,
		}, nil
	}
	_, err = model.CancelLike(*req.UserId, *req.VideoId)
	if err != nil {
		log.Println("取消点赞操作失败", err.Error())
		statuscode = 1
		statusmsg = "取消点赞操作失败"
		return &pb.DouyinFavoriteActionResponse{
			StatusCode: &statuscode,
			StatusMsg:  &statusmsg,
		}, nil
	}
	statuscode = 0
	statusmsg = "取消点赞成功"
	return &pb.DouyinFavoriteActionResponse{
		StatusCode: &statuscode,
		StatusMsg:  &statusmsg,
	}, nil
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

func main() {
	log.Println("正在启动FavoriteAction服务......")
	InitConfig()
	model.InitDB()
	model.InitRedis()
	//初始化grpc实例
	grpcServer := grpc.NewServer()

	//注册服务
	pb.RegisterFavoriteActionServer(grpcServer, new(FavoriteAction))

	//设置监听
	listen, err := net.Listen("tcp", ":12355")
	if err != nil {
		log.Println("注册服务启动监听失败")
	}
	defer listen.Close()

	//启动服务
	grpcServer.Serve(listen)
}
