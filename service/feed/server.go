package main

import (
	"context"
	"github.com/chaossat/tiktak/service/feed/model"
	"github.com/chaossat/tiktak/service/feed/pb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type Feed struct {
}

func (this *Feed) GetFeed(ctx context.Context, req *pb.DouyinFeedRequest) (*pb.DouyinFeedResponse, error) {
	//latest_time := req.GetLatestTime() //int64

	return nil, nil
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	InitConfig()
	model.InitDB()
	//初始化grpc实例
	grpcServer := grpc.NewServer()

	//注册服务
	pb.RegisterFeedServer(grpcServer, new(Feed))

	//设置监听
	listen, err := net.Listen("tcp", ":12346")
	if err != nil {
		log.Println("注册服务启动监听失败")
	}
	defer listen.Close()

	//启动服务
	grpcServer.Serve(listen)
}
