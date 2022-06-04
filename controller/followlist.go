package controller

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/chaossat/tiktak/middleware"

	"github.com/chaossat/tiktak/service/followlist/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func FollowListHandler(ctx *gin.Context) {
	//开启debug，观察性能瓶颈
	debugid, ok := <-DebugChan
	if ok {
		now := time.Now()
		log.Println("开始关注列表请求,操作ID:", debugid)
		defer log.Println("结束关注列表请求,操作ID:", debugid, "操作耗时：", time.Since(now))
	}

	//获得前端的请求
	token := ctx.Query("token")
	claims, err := middleware.CheckToken(token)
	if err != nil {
		log.Println("token错误", err)
		ctx.JSON(200, gin.H{
			"status_code": 1,
			"status_msg":  "token错误",
			"user_list":   nil,
		})
		return
	}
	user_id := claims.UserID
	//建立rpc客户端去请求server
	grpcConn, err := grpc.Dial(":12359", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("连接grpc服务失败")
		ctx.JSON(200, gin.H{
			"status_code": 1,
			"status_msg":  "连接grpc服务失败",
			"user_list":   nil,
		})
		return
	}
	defer grpcConn.Close()

	//初始化grpc客户端
	grpcClient := pb.NewFollowListClient(grpcConn)

	//创建并初始化registerrequest对象
	var req pb.DouyinRelationFollowListRequest
	req.UserId = &user_id
	req.Token = &token
	resp, err := grpcClient.GetFollowList(context.TODO(), &req)
	// log.Println("resp:", resp)
	if err != nil {
		log.Println(err.Error())
		log.Println("调用远程服务错误")
		ctx.JSON(200, gin.H{
			"status_code": 1,
			"status_msg":  "调用远程服务错误",
			"user_list":   nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
