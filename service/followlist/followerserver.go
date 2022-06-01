package main

import (
	"context"
	"fmt"
	"github.com/chaossat/tiktak/common"
	"github.com/chaossat/tiktak/db"
	"github.com/chaossat/tiktak/middleware"
	"github.com/chaossat/tiktak/service/followlist/pb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type Followlist struct {
	pb.UnimplementedFollowlistServer
}

func (followlist Followlist) GetFollowlist(ctx context.Context, request *pb.DouyinRelationFollowListRequest) (*pb.DouyinRelationFollowListResponse, error) {
	_, err := middleware.CheckToken(*request.Token)
	if err != nil {
		var code int32 = -1
		var msg string = "token认证失败" + err.Error()
		response := pb.DouyinRelationFollowListResponse{
			StatusCode: &code,
			StatusMsg:  &msg,
			UserList:   nil,
		}

		return &response, nil
	}
	userinf, err := db.UserInfoById(int(*request.UserId))
	if err != nil {
		var code int32 = -2
		var msg string = "查询用户信息失败！"
		response := pb.DouyinRelationFollowListResponse{
			StatusCode: &code,
			StatusMsg:  &msg,
			UserList:   nil,
		}
		return &response, nil
	}
	if userinf.ID == 0 {
		var code int32 = -3
		var msg string = "对应的id的用户不存在"
		response := pb.DouyinRelationFollowListResponse{
			StatusCode: &code,
			StatusMsg:  &msg,
			UserList:   nil,
		}
		return &response, nil
	}
	follows, err := db.FollowListByID(int(*request.UserId))
	if err != nil {
		var code int32 = -4
		var msg string = "查询用户的关注失败"
		response := pb.DouyinRelationFollowListResponse{
			StatusCode: &code,
			StatusMsg:  &msg,
			UserList:   nil,
		}
		return &response, err
	}

	var follows_ans = make([]*pb.User, len(follows))
	for i := 0; i < len(follows_ans); i++ {
		followCount, err := db.FollowCountByID(int(follows[i].ID))
		if err != nil {
			var code int32 = -5
			var msg string = "查询用户的关注数失败"
			response := pb.DouyinRelationFollowListResponse{
				StatusCode: &code,
				StatusMsg:  &msg,
				UserList:   nil,
			}
			return &response, err
		}
		followerCount, err := db.FollowerCountByID(int(follows[i].ID))
		if err != nil {
			var code int32 = -6
			var msg string = "查询用的粉丝数失败"
			response := pb.DouyinRelationFollowListResponse{
				StatusCode: &code,
				StatusMsg:  &msg,
				UserList:   nil,
			}
			return &response, err
		}
		isfollowdudu := db.IsFollow(userinf, *follows[i])

		follows_ans[i] = &pb.User{
			Id:            &follows[i].ID,
			Name:          &follows[i].Username,
			FollowCount:   &followCount,
			FollowerCount: &followerCount,
			IsFollow:      &isfollowdudu,
		}
	}
	var code int32 = 0
	var msg string = "验证成功！"
	response := &pb.DouyinRelationFollowListResponse{
		StatusCode: &code,
		StatusMsg:  &msg,
		UserList:   follows_ans,
	}
	return response, nil

}

func main() {
	InitConfig()
	common.InitDB()
	//初始化grpc实例
	grpcserver := grpc.NewServer()
	//注册服务
	pb.RegisterFollowlistServer(grpcserver, new(Followlist))
	//设置监听
	listen, err := net.Listen("tcp", ":17801")

	if err != nil {
		log.Println("注册服务启动监听失败")
	}
	defer listen.Close()

	//启动服务
	grpcserver.Serve(listen)
}

func InitConfig() {
	workDir, _ := os.Getwd() //获取当前工作的路径
	viper.SetConfigName("config.yml")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config/")
	fmt.Println("configPath:", workDir+"/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
