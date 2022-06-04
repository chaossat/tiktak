package controller

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/chaossat/tiktak/service/comment/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// CommentActionHandler:评论操作处理器函数
func CommentActionHandler(ctx *gin.Context) {
	//开启debug，观察性能瓶颈
	debugid, ok := <-DebugChan
	if ok {
		now := time.Now()
		log.Println("开始评论操作请求,操作ID:", debugid)
		defer log.Println("结束评论操作请求,操作ID:", debugid, "操作耗时：", time.Since(now))
	}

	userid, _ := strconv.ParseInt(ctx.Query("user_id"), 10, 64)
	token := ctx.Query("token")
	video_id, _ := strconv.ParseInt(ctx.Query("video_id"), 10, 64)
	action_type, _ := strconv.ParseInt(ctx.Query("action_type"), 10, 32)
	comment_text := ctx.Query("comment_text")
	comment_id, _ := strconv.ParseInt(ctx.Query("comment_id"), 10, 64)
	atype := int32(action_type)

	//连接grpc服务
	grpcConn, err := grpc.Dial(":10005", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("连接grpc服务失败")
		CommentActionResponse(ctx, -1, "Error Occoured!", nil)
		return
	}
	defer grpcConn.Close()
	//初始化grpc客户端
	grpcClient := pb.NewCommentActionClient(grpcConn)

	//创建并初始化对象
	var req pb.DouyinCommentActionRequest
	req.UserId = &userid
	req.Token = &token
	req.VideoId = &video_id
	req.ActionType = &atype
	req.CommentText = &comment_text
	req.CommentId = &comment_id

	//发起请求
	resp, err := grpcClient.CommentAction(context.TODO(), &req)
	// fmt.Println("commentAction resp:", resp)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("调用远程服务错误")
		CommentActionResponse(ctx, -2, "Error Occoured!", nil)
		return
	}
	if resp.GetStatusCode() != 0 {
		CommentActionResponse(ctx, -3, resp.GetStatusMsg(), nil)
		return
	}
	CommentActionResponse(ctx, 0, resp.GetStatusMsg(), resp.GetComment())
}

//CommentActionResponse：返回评论操作处理信息
func CommentActionResponse(ctx *gin.Context, code int, msg string, comment *pb.Comment) {
	jsonC := &pb.Comment{}
	if comment != nil {
		jsonC = comment
	}
	ctx.JSON(200, gin.H{
		"status_code": code,
		"status_msg":  msg,
		"comment":     *jsonC,
	})
}

// CommentListHandler:评论列表处理器函数
func CommentListHandler(ctx *gin.Context) {
	//开启debug，观察性能瓶颈
	debugid, ok := <-DebugChan
	if ok {
		now := time.Now()
		log.Println("开始评论列表请求,操作ID:", debugid)
		defer log.Println("结束评论列表请求,操作ID:", debugid, "操作耗时：", time.Since(now))
	}

	token := ctx.Query("token")
	video_id, _ := strconv.ParseInt(ctx.Query("video_id"), 10, 64)
	//连接grpc服务
	grpcConn, err := grpc.Dial(":10005", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("连接grpc服务失败")
		CommentListResponse(ctx, -1, "Error Occoured!", nil)
		return
	}
	defer grpcConn.Close()
	//初始化grpc客户端
	grpcClient := pb.NewCommentListClient(grpcConn)
	//初始化请求对象
	var req = pb.DouyinCommentListRequest{}
	req.Token = &token
	req.VideoId = &video_id
	resp, err := grpcClient.CommentList(context.TODO(), &req)
	// fmt.Println("commentList resp:", resp)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("调用远程服务错误")
		CommentListResponse(ctx, -2, "Error Occoured!", nil)
		return
	}
	if resp.GetStatusCode() != 0 {
		CommentListResponse(ctx, -3, resp.GetStatusMsg(), nil)
		return
	}
	ctx.JSON(200, resp)
}

//CommentListResponse：返回评论列表处理信息
func CommentListResponse(ctx *gin.Context, code int, msg string, comments []*pb.Comment) {
	ctx.JSON(200, gin.H{
		"StatusCode":  code,
		"StatusMsg":   msg,
		"CommentList": comments,
	})
}
