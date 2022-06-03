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
	"github.com/chaossat/tiktak/service/login/pb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	log.Println("正在启动Login服务......")
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

type LoginHandler struct {
	pb.UnimplementedLoginServer
}

//Login:grpc的Login接口
func (loginHandler *LoginHandler) Login(ctx context.Context, loginRequest *pb.DouyinUserLoginRequest) (*pb.DouyinUserLoginResponse, error) {
	username := loginRequest.GetUsername()
	password := loginRequest.GetPassword()
	ok, err := db.UserLogin(username, password)
	if err != nil {
		fmt.Println(err.Error())
		return LoginResponse(-1, "Error Occoured!", 0, ""), nil
	}
	if !ok {
		return LoginResponse(-2, "Invalid Verify Imformation!", 0, ""), nil
	}
	err, userinfo := db.UserInfoByName(username)
	if err != nil {
		fmt.Println(err.Error())
		return LoginResponse(-3, "Error Occoured!", 0, ""), nil
	}
	token, err := middleware.CreateToken(userinfo.ID)
	if err != nil {
		fmt.Println(err.Error())
		return LoginResponse(-4, "Error Occoured!", 0, ""), nil
	}
	return LoginResponse(0, "Login Succeed!", int(userinfo.ID), token), nil
}

//LoginResponse:返回登录处理信息
func LoginResponse(code int, msg string, user_id int, token string) *pb.DouyinUserLoginResponse {
	codeResponse := int32(code)
	userID := int64(user_id)
	return &pb.DouyinUserLoginResponse{
		StatusCode: &codeResponse,
		StatusMsg:  &msg,
		UserId:     &userID,
		Token:      &token,
	}
}
