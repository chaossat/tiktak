package main

import (
	"context"
	"github.com/chaossat/tiktak/middleware"
	"github.com/chaossat/tiktak/service/register/model"
	"github.com/chaossat/tiktak/service/register/pb"
	"github.com/chaossat/tiktak/util"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type Register struct{}

func (this *Register) Register(ctx context.Context, req *pb.DouyinUserRegisterRequest) (*pb.DouyinUserRegisterResponse, error) {
	username := req.GetUsername()
	password := req.GetPassword()
	password_hash := util.Sha1([]byte(password))
	user, err := model.SaveUser(username, password_hash)
	var statuscode int32
	var statusmsg string
	var userid int64
	var token string
	if err != nil {
		if err.Error() == "用户名已存在" {
			log.Println("发现用户名已存在错误...")
			statuscode = int32(1)
			statusmsg = "用户名已存在"
			userid = int64(0)
			token = ""
			resp := pb.DouyinUserRegisterResponse{
				StatusCode: &statuscode,
				StatusMsg:  &statusmsg,
				UserId:     &userid,
				Token:      &token,
			}
			log.Println("用户已存在 resp:", resp)
			return &resp, nil
		}
		statusmsg = "数据库存储失败"
		return &pb.DouyinUserRegisterResponse{
			StatusCode: &statuscode,
			StatusMsg:  &statusmsg,
			UserId:     &userid,
			Token:      &token,
		}, nil
	}
	token, err = middleware.CreateToken(username)
	if err != nil {
		log.Println("生成token失败", err.Error())
		statuscode = 1
		statusmsg = "token生成失败"
		userid = 0
		token = ""
		return &pb.DouyinUserRegisterResponse{
			StatusCode: &statuscode,
			StatusMsg:  &statusmsg,
			UserId:     &userid,
			Token:      &token,
		}, nil
	}
	//log.Println("token:", token)
	statuscode = 0
	statusmsg = "注册成功"
	userid = user.ID
	resp := pb.DouyinUserRegisterResponse{
		StatusCode: &statuscode,
		StatusMsg:  &statusmsg,
		UserId:     &userid,
		Token:      &token,
	}
	return &resp, nil
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
	pb.RegisterRegisterServer(grpcServer, new(Register))

	//设置监听
	listen, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Println("注册服务启动监听失败")
	}
	defer listen.Close()

	//启动服务
	grpcServer.Serve(listen)

}
