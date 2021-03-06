package controller

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/chaossat/tiktak/middleware"
	"github.com/chaossat/tiktak/service/followaction/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func FollowActionHandler(ctx *gin.Context) {
	//开启debug，观察性能瓶颈
	debugid, ok := <-DebugChan
	if ok {
		now := time.Now()
		log.Println("开始关注操作请求,操作ID:", debugid)
		defer log.Println("结束关注操作请求,操作ID:", debugid, "操作耗时：", time.Since(now))
	}

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
	grpcConn, err := grpc.Dial(":12358", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
	grpcClient := pb.NewFollowActionClient(grpcConn)

	//创建并初始化registerrequest对象
	var req pb.DouyinRelationActionRequest
	req.UserId = &user_id
	req.Token = &token
	req.ToUserId = &touseridint
	req.ActionType = &actiontypeint
	resp, err := grpcClient.FollowAction(context.TODO(), &req)
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
