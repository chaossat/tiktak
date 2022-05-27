package controller

import (
	"context"
	"fmt"
	"strconv"

	"github.com/chaossat/tiktak/service/userinf/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 根据GET请求的id和token获取用户参数
func UserInfHandler(ctx *gin.Context) {
	token := ctx.Query("token")
	user_id, err := strconv.Atoi(ctx.Query("user_id"))
	uid := int64(user_id)
	if err != nil {
		fmt.Println(err.Error())
		UserinfoResponse(ctx, -1, "Error Occured!", pb.User{})
	}
	//连接grpc服务
	grpcConn, err := grpc.Dial(":10003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("连接grpc服务失败")
		UserinfoResponse(ctx, -1, "Error Occoured!", pb.User{})
		return
	}
	defer grpcConn.Close()
	//初始化grpc客户端
	grpcClient := pb.NewUserInfClient(grpcConn)

	//创建并初始化UserinfRequest对象
	var req pb.DouyinUserRequest
	req.Token = &token
	req.UserId = &uid

	resp, err := grpcClient.GetUserinf(context.TODO(), &req)
	fmt.Println("userinfo resp:", resp)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("调用远程服务错误")
		UserinfoResponse(ctx, -2, "Error Occoured!", pb.User{})
		return
	}
	if resp.GetStatusCode() != 0 {
		UserinfoResponse(ctx, -3, resp.GetStatusMsg(), pb.User{})
		return
	}
	UserinfoResponse(ctx, resp.GetStatusCode(), resp.GetStatusMsg(), *resp.GetUser())
}

//UserinfoResponse:返回发布列表处理信息
func UserinfoResponse(ctx *gin.Context, code int32, msg string, user pb.User) {
	ctx.JSON(200, gin.H{
		"status_code": code,
		"status_msg":  msg,
		"user":        user,
	})
}
