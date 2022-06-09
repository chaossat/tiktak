package controller

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/chaossat/tiktak/service/feed/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Feed(ctx *gin.Context) {
	//开启debug，观察性能瓶颈
	debugid, ok := <-DebugChan
	if ok {
		now := time.Now()
		log.Println("开始Feed请求,操作ID:", debugid)
		defer log.Println("结束Feed请求,操作ID:", debugid, "操作耗时：", time.Since(now))
	}

	latest_time := ctx.Query("latest_time")
	latesttime, err := strconv.ParseInt(latest_time, 10, 64)
	if err != nil {
		log.Println("请求时间戳错误")
		ctx.JSON(200, gin.H{
			"status_code": 1,
			"status_msg":  "请求时间戳错误",
		})
		return
	}
	token := ctx.Query("token")
	log.Println("token:", token)
	//连接grpc服务
	grpcConn, err := grpc.Dial(":12346", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("连接grpc服务失败")
		ctx.JSON(200, gin.H{
			"status_code": 1,
			"status_msg":  "连接grpc服务失败",
		})
		return
	}
	defer grpcConn.Close()

	//初始化grpc客户端
	grpcClient := pb.NewFeedClient(grpcConn)

	//创建并初始化registerrequest对象
	var req pb.DouyinFeedRequest
	req.LatestTime = &latesttime
	req.Token = &token
	resp, err := grpcClient.GetFeed(context.TODO(), &req)
	if err != nil {
		log.Println(err.Error())
		log.Println("调用远程服务错误")
		ctx.JSON(200, gin.H{
			"status_code": 1,
			"status_msg":  "调用远程服务错误",
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
