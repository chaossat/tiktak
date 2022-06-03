package controller

import (
	"context"
	"fmt"
	"strconv"

	"github.com/chaossat/tiktak/service/followerlist/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Followerlist(ctx *gin.Context) {
	token := ctx.Query("token")
	user_id, err := strconv.Atoi(ctx.Query("user_id"))
	uid := int64(user_id)
	if err != nil {
		fmt.Println(err.Error())
		FollowerResponse(ctx, -1, "Error Occured!", nil)
	}
	//连接grpc服务
	grpcConn, err := grpc.Dial(":17800", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("连接grpc服务失败")
		FollowerResponse(ctx, -2, "Error Occoured!", nil)
		return
	}
	defer grpcConn.Close()
	//初始化grpc客户端
	grpcClient := pb.NewFollowerlistClient(grpcConn)
	var request pb.DouyinRelationFollowerListRequest
	request.Token = &token
	request.UserId = &uid
	response, err := grpcClient.GetFollowerlist(context.TODO(), &request)
	fmt.Println("publist resp:", response)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("调用远程服务错误")
		FollowerResponse(ctx, -3, "Error Occoured!", nil)
		return
	}
	if response.GetStatusCode() != 0 {
		FollowerResponse(ctx, -4, response.GetStatusMsg(), nil)
		return
	}

	FollowerResponse(ctx, response.GetStatusCode(), response.GetStatusMsg(), response.GetUserList())
}

//FollowertResponse:返回处理信息
func FollowerResponse(ctx *gin.Context, code int32, msg string, user_list []*pb.User) {
	ctx.JSON(200, gin.H{
		"status_code": code,
		"status_msg":  msg,
		"user_list":   user_list,
	})
}
