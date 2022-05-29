package controller

import (
	"context"
	"github.com/chaossat/tiktak/middleware"
	"github.com/chaossat/tiktak/service/favoriteaction/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"strconv"
)

func FavoriteAction(ctx *gin.Context) {
	//userid := ctx.Query("user_id")
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
	//user_id, err := strconv.ParseInt(userid, 10, 64)
	//if err != nil {
	//	log.Println("userid错误", err)
	//	ctx.JSON(200, gin.H{
	//		"status_code": 1,
	//		"status_msg":  "userid错误",
	//	})
	//	return
	//}
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
