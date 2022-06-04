package controller

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/chaossat/tiktak/service/favoritelist/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func FavoriteList(ctx *gin.Context) {
	//开启debug，观察性能瓶颈
	debugid, ok := <-DebugChan
	if ok {
		now := time.Now()
		log.Println("开始点赞列表请求,操作ID:", debugid)
		defer log.Println("结束点赞列表请求,操作ID:", debugid, "操作耗时：", time.Since(now))
	}

	userid := ctx.Query("user_id")
	token := ctx.Query("token")
	log.Println("userid:", userid)
	log.Println("token:", token)
	user_id, err := strconv.ParseInt(userid, 10, 64)
	if err != nil {
		log.Println("userid错误")
		ctx.JSON(200, gin.H{
			"status_code": 1,
			"status_msg":  "userid错误",
		})
		return
	}
	grpcConn, err := grpc.Dial(":12350", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
	grpcClient := pb.NewFavoriteListClient(grpcConn)

	//创建并初始化registerrequest对象
	var req pb.DouyinFavoriteListRequest
	req.UserId = &user_id
	req.Token = &token
	resp, err := grpcClient.GetFavoriteList(context.TODO(), &req)
	// log.Println("resp:", resp)
	if err != nil {
		log.Println(err.Error())
		log.Println("点赞过的视频调用远程服务错误")
		ctx.JSON(200, gin.H{
			"status_code": 1,
			"status_msg":  "调用远程服务错误",
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
