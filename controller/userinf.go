package controller

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/chaossat/tiktak/db"
	feedmodel "github.com/chaossat/tiktak/service/feed/model"
	"github.com/chaossat/tiktak/service/userinf/pb"
	"github.com/chaossat/tiktak/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 根据GET请求的id和token获取用户参数
func UserInfHandler(ctx *gin.Context) {
	//开启debug，观察性能瓶颈
	debugid, ok := <-DebugChan
	if ok {
		now := time.Now()
		log.Println("开始用户信息请求,操作ID:", debugid)
		defer log.Println("结束用户信息请求,操作ID:", debugid, "操作耗时：", time.Since(now))
	}

	token := ctx.Query("token")
	user_id, err := strconv.Atoi(ctx.Query("user_id"))
	uid := int64(user_id)
	if err != nil {
		fmt.Println(err.Error())
		UserinfoResponse(ctx, -1, "Error Occured!", pb.User{})
	}
	//连接grpc服务
	grpcConn, err := grpc.Dial(":10003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("连接grpc服务失败")
		UserinfoResponse(ctx, -1, "Error Occoured!", pb.User{})
		return
	}
	defer grpcConn.Close()
	//初始化grpc客户端
	grpcClient := pb.NewUserInfClient(grpcConn)

	//创建并初始化UserinfRequest对象
	var req pb.DouyinUserRequest
	req.Token = &token
	req.UserId = &uid

	resp, err := grpcClient.GetUserinf(context.TODO(), &req)
	// fmt.Println("userinfo resp:", resp)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("调用远程服务错误")
		UserinfoResponse(ctx, -2, "Error Occoured!", pb.User{})
		return
	}
	if resp.GetStatusCode() != 0 {
		UserinfoResponse(ctx, -3, resp.GetStatusMsg(), pb.User{})
		return
	}
	// videoCount, err := db.VideoCountByID(user_id)
	// if err != nil {
	// 	fmt.Println("获取视频数失败")
	// 	UserinfoResponse(ctx, -4, "Error Occoured!", pb.User{})
	// 	return
	// }
	favoriteCount, err := db.FavoriteCountByID(user_id)
	if err != nil {
		fmt.Println("获取点赞数失败:", err.Error())
		UserinfoResponse(ctx, -5, "Error Occoured!", pb.User{})
		return
	}
	total_favorited, err := TotalFavorited(user_id)
	if err != nil {
		fmt.Println("获取被赞数失败：", err.Error())
		UserinfoResponse(ctx, -6, "Error Occoured!", pb.User{})
		return
	}
	type dtoUser struct {
		pb.User
		Total_favorited  int64  `json:"total_favorited"`
		Favorite_count   int64  `json:"favorite_count"`
		Avatar           string `json:"avatar"`
		Signature        string `json:"signature"`
		Background_image string `json:"background_image"`
	}
	ip := util.GetIP()
	dtouser := dtoUser{
		User:             *resp.GetUser(),
		Total_favorited:  total_favorited,
		Favorite_count:   int64(favoriteCount),
		Avatar:           ip + viper.GetString("server.port") + "/" + "static/avatar.png",
		Signature:        "第三届字节跳动青训营-后端专场  6824Nil",
		Background_image: ip + viper.GetString("server.port") + "/" + "static/background.jpg",
	}
	ctx.JSON(200, gin.H{
		"status_code": resp.GetStatusCode(),
		"status_msg":  resp.GetStatusMsg(),
		"user":        dtouser,
	})
}

//UserinfoResponse:返回发布列表处理信息
func UserinfoResponse(ctx *gin.Context, code int32, msg string, user pb.User) {
	ctx.JSON(200, gin.H{
		"status_code": code,
		"status_msg":  msg,
		"user":        user,
	})
}

//TotalFavorited：获取用户总获赞数，根据客户端的调整临时追加，效率很低
func TotalFavorited(uid int) (res int64, err error) {
	videos, err := db.VideoedByID(uid)
	if err != nil {
		if err.Error() == "record not found" {
			return 0, nil
		}
		return 0, err
	}
	for _, video := range videos {
		favoritecnt, err := feedmodel.GetFavoriteCount(video.ID)
		if err != nil {
			return 0, err
		}
		res += favoritecnt
	}
	return res, nil
}
