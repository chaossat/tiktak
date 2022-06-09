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
	"github.com/chaossat/tiktak/service/userinf/pb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type UserInf struct {
	pb.UnimplementedUserInfServer
}

/**
获取用户信息服务
*/
var errorID = int64(0)
var errorName = "Null"
var errorFollow = false
var errorUser = pb.User{
	Id:       &errorID,
	Name:     &errorName,
	IsFollow: &errorFollow,
}

func (u *UserInf) GetUserinf(context context.Context, request *pb.DouyinUserRequest) (*pb.DouyinUserResponse, error) {
	uid := request.UserId
	token := request.Token
	_, err := middleware.CheckToken(*token)
	if err != nil {
		var code int32 = -1
		var msg string = "token认证失败" + err.Error()
		response := pb.DouyinUserResponse{
			StatusCode: &code,
			StatusMsg:  &msg,
			User:       &errorUser,
		}
		return &response, nil
	}
	// 查询数据库获取用户信息
	inf, err := db.UserInfoById(int(*uid))
	if err != nil || inf.ID == 0 {
		var code int32 = -2
		var msg string = "查询用户信息失败"
		response := pb.DouyinUserResponse{
			StatusCode: &code,
			StatusMsg:  &msg,
			User:       &errorUser,
		}
		return &response, nil
	}
	// 用户信息结构体
	var isFollow = false
	followCount, err := db.FollowCountByID(int(*uid))
	if err != nil || inf.ID == 0 {
		var code int32 = -3
		var msg string = "查询用户关注数失败"
		response := pb.DouyinUserResponse{
			StatusCode: &code,
			StatusMsg:  &msg,
			User:       &errorUser,
		}
		return &response, nil
	}
	followerCount, err := db.FollowerCountByID(int(*uid))
	if err != nil || inf.ID == 0 {
		var code int32 = -4
		var msg string = "查询用户粉丝数失败"
		response := pb.DouyinUserResponse{
			StatusCode: &code,
			StatusMsg:  &msg,
			User:       &errorUser,
		}
		return &response, nil
	}
	var user = pb.User{
		Id:            uid,
		Name:          &inf.Username,
		FollowCount:   &followCount,
		FollowerCount: &followerCount,
		IsFollow:      &isFollow,
	}
	var code int32 = 0
	var msg string = "成功！"
	return &pb.DouyinUserResponse{
		StatusCode: &code,
		StatusMsg:  &msg,
		User:       &user,
	}, nil
}

func main() {
	log.Println("正在启动Userinf服务......")
	InitConfig()
	common.InitDB()
	//初始化grpc实例
	grpcServer := grpc.NewServer()

	//注册服务
	pb.RegisterUserInfServer(grpcServer, new(UserInf))

	//设置监听
	listen, err := net.Listen("tcp", ":10003")
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
