package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/CodyGuo/godaemon"
	"github.com/chaossat/tiktak/common"
	"github.com/chaossat/tiktak/db"
	"github.com/chaossat/tiktak/middleware"
	"github.com/chaossat/tiktak/service/followlist/pb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type FollowList struct {
	pb.UnimplementedFollowListServer
}

//服务
func (f *FollowList) GetFollowList(ctx context.Context, req *pb.DouyinRelationFollowListRequest) (*pb.DouyinRelationFollowListResponse, error) {
	// token验证
	_, err := middleware.CheckToken(*req.Token)
	if err != nil {
		var code int32 = -1
		var msg string = "验证失败！"
		return &pb.DouyinRelationFollowListResponse{
			StatusCode: &code,
			StatusMsg:  &msg,
			UserList:   nil,
		}, nil
	}
	//验证通过,获取用户信息
	var (
		statuscode int32
		statusmsg  string
		// userlist   []*model.User
	)
	userinf, err := db.UserInfoById(int(*req.UserId))
	if err != nil || userinf.ID == 0 {
		statuscode = 1
		statusmsg = "验证user_id失败"
		return &pb.DouyinRelationFollowListResponse{
			StatusCode: &statuscode,
			StatusMsg:  &statusmsg,
			UserList:   nil,
		}, err
	}
	//获取关注
	follows, err := db.FollowListByID(int(*req.UserId))
	if err != nil {
		statuscode = int32(-4)
		statusmsg = "查询用户粉丝失败！"
		response := pb.DouyinRelationFollowListResponse{
			StatusCode: &statuscode,
			StatusMsg:  &statusmsg,
			UserList:   nil,
		}
		return &response, err
	}
	log.Println("开始查询")
	var follows_ans = make([]*pb.User, len(follows))
	for i := 0; i < len(follows); i++ {
		followerCount, err := db.FollowerCountByID(int(follows[i].ID))
		if err != nil {
			statuscode = -6
			statusmsg = "查询粉丝数量失败！"
			return &pb.DouyinRelationFollowListResponse{
				StatusCode: &statuscode,
				StatusMsg:  &statusmsg,
				UserList:   nil,
			}, err
		}
		followCount, err := db.FollowCountByID(int(follows[i].ID))
		if err != nil {
			statuscode = -6
			statusmsg = "查询关注数量失败！"
			return &pb.DouyinRelationFollowListResponse{
				StatusCode: &statuscode,
				StatusMsg:  &statusmsg,
				UserList:   nil,
			}, err
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
	statuscode = 0
	statusmsg = "验证成功！"
	return &pb.DouyinRelationFollowListResponse{
		StatusCode: &statuscode,
		StatusMsg:  &statusmsg,
		UserList:   follows_ans,
	}, nil
}
func main() {
	log.Println("正在启动FollowList服务......")
	InitConfig()
	common.InitDB()
	common.InitRedis()
	//初始化grpc实例
	grpcServer := grpc.NewServer()
	//注册服务
	pb.RegisterFollowListServer(grpcServer, new(FollowList))
	//设置监听
	listen, err := net.Listen("tcp", ":12359")
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
