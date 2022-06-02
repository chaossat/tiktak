package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/chaossat/tiktak/common"
	"github.com/chaossat/tiktak/db"
	"github.com/chaossat/tiktak/middleware"
	"github.com/chaossat/tiktak/service/followerlist/pb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type Followerlist struct {
	pb.UnimplementedFollowerlistServer
}

func (followerlist Followerlist) GetFollowerlist(ctx context.Context, request *pb.DouyinRelationFollowerListRequest) (*pb.DouyinRelationFollowerListResponse, error) {

	_, err := middleware.CheckToken(*request.Token)
	if err != nil {
		var code int32 = -1
		var msg string = "token认证失败" + err.Error()
		response := pb.DouyinRelationFollowerListResponse{
			StatusCode: &code,
			StatusMsg:  &msg,
			UserList:   nil,
		}
		return &response, nil
	}
	userinf, err := db.UserInfoById(int(*request.UserId))

	if err != nil {
		var code int32 = -2
		var msg string = "查询用户信息失败!"
		response := pb.DouyinRelationFollowerListResponse{
			StatusCode: &code,
			StatusMsg:  &msg,
			UserList:   nil,
		}
		return &response, nil
	}

	if userinf.ID == 0 {
		var code int32 = -3
		var msg string = "对应id的用户不存在!"
		response := pb.DouyinRelationFollowerListResponse{
			StatusCode: &code,
			StatusMsg:  &msg,
			UserList:   nil,
		}
		return &response, nil
	}
	followers, err := db.FollowerListByID(int(*request.UserId))
	if err != nil {
		var code int32 = -4
		var msg string = "查询用户粉丝失败!"
		response := pb.DouyinRelationFollowerListResponse{
			StatusCode: &code,
			StatusMsg:  &msg,
			UserList:   nil,
		}
		return &response, nil
	}
	var followers_ans = make([]*pb.User, len(followers))

	for i := 0; i < len(followers); i++ {
		followerCount, err := db.FollowerCountByID(int(followers[i].ID))
		if err != nil {

			var code int32 = -5
			var msg string = "查询粉丝数量失败！"

			return &pb.DouyinRelationFollowerListResponse{
				StatusCode: &code,
				StatusMsg:  &msg,
				UserList:   nil,
			}, nil
		}
		followCount, err := db.FollowCountByID(int(followers[i].ID))
		if err != nil {
			var code int32 = -6
			var msg string = "查询关注数量失败!"
			return &pb.DouyinRelationFollowerListResponse{
				StatusCode: &code,
				StatusMsg:  &msg,
				UserList:   nil,
			}, nil
		}
		isfollowdudu := db.IsFollow(userinf, *followers[i])
		followers_ans[i] = &pb.User{
			Id:            &followers[i].ID,
			Name:          &followers[i].Username,
			FollowCount:   &followCount,
			FollowerCount: &followerCount,
			IsFollow:      &isfollowdudu,
		}
	}
	var code int32 = 0
	var msg string = "验证成功!"
	response := &pb.DouyinRelationFollowerListResponse{
		StatusCode: &code,
		StatusMsg:  &msg,
		UserList:   followers_ans,
	}
	return response, err
}

func main() {
	log.Println("正在启动FollowerList服务......")
	InitConfig()
	common.InitDB()
	//初始化grpc实例
	grpcServer := grpc.NewServer()
	//注册服务
	pb.RegisterFollowerlistServer(grpcServer, new(Followerlist))
	//设置监听
	listen, err := net.Listen("tcp", ":17800")
	if err != nil {
		log.Println("注册服务启动监听失败")
	}
	defer listen.Close()

	//启动服务
	grpcServer.Serve(listen)
}

func InitConfig() {
	workDir, _ := os.Getwd() //获取当前工作路径，非文件路径，以终端显示路径为准
	viper.SetConfigName("config.yml")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	fmt.Println("configPath:", workDir+"/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
