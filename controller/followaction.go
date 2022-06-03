package controller

import (
	"context"
	"log"
	"strconv"

	"github.com/chaossat/tiktak/middleware"
	"github.com/chaossat/tiktak/service/followaction/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func FollowActionHandler(ctx *gin.Context) {
	token := ctx.Query("token")
	touserid := ctx.Query("to_user_id")
	actiontype := ctx.Query("action_type")
	claims, err := middleware.CheckToken(token)
	if err != nil {
		log.Println("token错误", err)
		ctx.JSON(200, gin.H{
			"status_code": 1,
			"status_msg":  "token错误",
		})
		return
	}
	user_id := claims.UserID
	touseridint, _ := strconv.ParseInt(touserid, 10, 64)
	tmp, _ := strconv.ParseInt(actiontype, 10, 64)
	actiontypeint := int32(tmp)
	// log.Println("连接grpc服务")
	grpcConn, err := grpc.Dial(":12358", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("连接grpc服务失败")
		ctx.JSON(200, gin.H{
			"status_code": 1,
			"status_msg":  "连接grpc服务失败",
		})
		return
	}
	// log.Println("连接grpc服务成功")
	defer grpcConn.Close()

	//初始化grpc客户端
	grpcClient := pb.NewFollowActionClient(grpcConn)

	//创建并初始化registerrequest对象
	var req pb.DouyinRelationActionRequest
	req.UserId = &user_id
	req.Token = &token
	req.ToUserId = &touseridint
	req.ActionType = &actiontypeint
	log.Println("req:", *req.UserId, *req.ActionType, *req.ToUserId)
	resp, err := grpcClient.FollowAction(context.TODO(), &req)
	log.Println("resp:", resp)
	if err != nil {
		log.Println(err.Error())
		log.Println("调用远程服务错误")
		ctx.JSON(200, gin.H{
			"status_code": 1,
			"status_msg":  "调用远程服务错误",
		})
		return
	}
	ctx.JSON(200, resp)
}
