package controller

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/chaossat/tiktak/middleware"
	"github.com/chaossat/tiktak/service/favoriteaction/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func FavoriteAction(ctx *gin.Context) {
	//开启debug，观察性能瓶颈
	debugid, ok := <-DebugChan
	if ok {
		now := time.Now()
		log.Println("开始点赞操作请求,操作ID:", debugid)
		defer log.Println("结束点赞操作请求,操作ID:", debugid, "操作耗时：", time.Since(now))
	}

	token := ctx.Query("token")
	videoid := ctx.Query("video_id")
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
	video_id, err := strconv.ParseInt(videoid, 10, 64)
	if err != nil {
		log.Println("videoid错误")
		ctx.JSON(200, gin.H{
			"status_code": 1,
			"status_msg":  "videoid错误",
		})
		return
	}
	tmp, err := strconv.ParseInt(actiontype, 10, 32)
	action_type := int32(tmp)
	if err != nil {
		log.Println("videoid错误")
		ctx.JSON(200, gin.H{
			"status_code": 1,
			"status_msg":  "videoid错误",
		})
		return
	}
	grpcConn, err := grpc.Dial(":12355", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
	grpcClient := pb.NewFavoriteActionClient(grpcConn)

	//创建并初始化registerrequest对象
	var req pb.DouyinFavoriteActionRequest
	req.UserId = &user_id
	req.Token = &token
	req.VideoId = &video_id
	req.ActionType = &action_type
	resp, err := grpcClient.GetFavoriteAction(context.TODO(), &req)
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
