package controller

import (
	"context"
	"github.com/chaossat/tiktak/service/register/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

//注册并生成token
func Register(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	//连接grpc服务
	grpcConn, err := grpc.Dial(":12345", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("连接grpc服务失败")
	}
	defer grpcConn.Close()

	//初始化grpc客户端
	grpcClient := pb.NewRegisterClient(grpcConn)

	//创建并初始化registerrequest对象
	var register pb.DouyinUserRegisterRequest
	register.Username = &username
	register.Password = &password
	resp, err := grpcClient.Register(context.TODO(), &register)
	log.Println("resp:", resp)
	if err != nil {
		log.Println(err.Error())
		log.Println("调用远程服务错误")
	}
	ctx.JSON(http.StatusOK, resp)
}
