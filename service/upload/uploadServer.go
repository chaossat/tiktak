package main

import (
	"log"
	"net"
	"os"

	"github.com/chaossat/tiktak/common"
	"github.com/chaossat/tiktak/service/login/pb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	InitConfig()
	common.InitDB()
	//初始化grpc实例
	grpcServer := grpc.NewServer()

	//注册服务
	pb.RegisterLoginServer(grpcServer, new(LoginHandler))

	//设置监听
	listen, err := net.Listen("tcp", ":10001")
	if err != nil {
		log.Println("注册服务启动监听失败")
	}
	defer listen.Close()

	//启动服务
	grpcServer.Serve(listen)
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config.yml")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
