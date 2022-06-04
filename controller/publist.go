package controller

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/chaossat/tiktak/service/publist/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//VideoListHandler:视频列表接口
func PubListHandler(ctx *gin.Context) {
	//开启debug，观察性能瓶颈
	debugid, ok := <-DebugChan
	if ok {
		now := time.Now()
		log.Println("开始视频列表请求,操作ID:", debugid)
		defer log.Println("结束视频列表请求,操作ID:", debugid, "操作耗时：", time.Since(now))
	}

	token := ctx.Query("token")
	user_id, err := strconv.Atoi(ctx.Query("user_id"))
	uid := int64(user_id)
	if err != nil {
		fmt.Println(err.Error())
		PublistResponse(ctx, -1, "Error Occured!", nil)
	}
	//连接grpc服务
	grpcConn, err := grpc.Dial(":10002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("连接grpc服务失败")
		PublistResponse(ctx, -1, "Error Occoured!", nil)
		return
	}
	defer grpcConn.Close()
	//初始化grpc客户端
	grpcClient := pb.NewPublishClient(grpcConn)

	//创建并初始化PublistRequest对象
	var req pb.DouyinPublishListRequest
	req.Token = &token
	req.UserId = &uid

	resp, err := grpcClient.PublishVideo(context.TODO(), &req)
	// fmt.Println("publist resp:", resp)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("调用远程服务错误")
		PublistResponse(ctx, -2, "Error Occoured!", nil)
		return
	}
	if resp.GetStatusCode() != 0 {
		PublistResponse(ctx, -3, resp.GetStatusMsg(), nil)
		return
	}
	PublistResponse(ctx, resp.GetStatusCode(), resp.GetStatusMsg(), resp.GetVideoList())
}

//PublistResponse:返回发布列表处理信息
func PublistResponse(ctx *gin.Context, code int32, msg string, video_list []*pb.Video) {
	ctx.JSON(200, gin.H{
		"status_code": code,
		"status_msg":  msg,
		"video_list":  video_list,
	})
}
