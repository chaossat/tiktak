package controller

import (
	"context"
	"fmt"
	"github.com/chaossat/tiktak/service/followlist/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strconv"
)

func Followlist(ctx *gin.Context) {
	token := ctx.Query("token")
	user_id, err := strconv.Atoi(ctx.Query("user_id"))
	uid := int64(user_id)
	if err != nil {
		fmt.Println(err.Error())
		FollowResponse(ctx, -1, "Error Occured!", nil)
	}
	//连接grpc服务
	grpcConn, err := grpc.Dial(":17801", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("连接grpc服务失败")
		FollowResponse(ctx, -1, "Error Occoured!", nil)
		return
	}
	defer grpcConn.Close()
	//初始化grpc客户端
	grpcClient := pb.NewFollowlistClient(grpcConn)
	var request pb.DouyinRelationFollowListRequest

	request.Token = &token
	request.UserId = &uid

	response, err := grpcClient.GetFollowlist(context.TODO(), &request)

	fmt.Println("publist resp:", response)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("调用远程服务错误")
		FollowResponse(ctx, -2, "Error Occoured!", nil)
		return
	}
	if response.GetStatusCode() != 0 {
		FollowerResponse(ctx, -3, response.GetStatusMsg(), nil)
		return
	}
	FollowResponse(ctx, response.GetStatusCode(), response.GetStatusMsg(), response.GetUserList())
}

func FollowResponse(ctx *gin.Context, code int32, msg string, user_list []*pb.User) {
	ctx.JSON(200, gin.H{
		"status_code": code,
		"status_msg":  msg,
		"user_list":   user_list,
	})
}
