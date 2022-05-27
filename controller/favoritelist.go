package controller

import (
	"context"
	"github.com/chaossat/tiktak/service/favoritelist/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"strconv"
)

func FavoriteList(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, resp)
}
