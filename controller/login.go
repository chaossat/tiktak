package controller

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chaossat/tiktak/service/login/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//LoginHandler:登录接口处理器
func LoginHandler(ctx *gin.Context) {
	//开启debug，观察性能瓶颈
	debugid, ok := <-DebugChan
	if ok {
		now := time.Now()
		log.Println("开始登录请求,操作ID:", debugid)
		defer log.Println("结束登录请求,操作ID:", debugid, "操作耗时：", time.Since(now))
	}

	username := ctx.Query("username")
	password := ctx.Query("password")
	//连接grpc服务
	grpcConn, err := grpc.Dial(":10001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("连接grpc服务失败")
		LoginResponse(ctx, -1, "Error Occoured!", 0, "")
		return
	}
	defer grpcConn.Close()
	//初始化grpc客户端
	grpcClient := pb.NewLoginClient(grpcConn)
	//创建并初始化LoginRequest对象
	var req pb.DouyinUserLoginRequest
	req.Username = &username
	req.Password = &password

	resp, err := grpcClient.Login(context.TODO(), &req)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("调用远程服务错误")
		LoginResponse(ctx, -2, "Error Occoured!", 0, "")
		return
	}
	if resp.GetStatusCode() != 0 {
		LoginResponse(ctx, -3, resp.GetStatusMsg(), 0, "")
		return
	}
	LoginResponse(ctx, 0, "Login Succeed!", int(resp.GetUserId()), resp.GetToken())
}

//LoginResponse:返回登录处理信息
func LoginResponse(ctx *gin.Context, code int, msg string, user_id int, token string) {
	ctx.JSON(200, gin.H{
		"status_code": code,
		"status_msg":  msg,
		"user_id":     user_id,
		"token":       token,
	})
}
