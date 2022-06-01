package followlist

import (
	"context"
	"fmt"
	"github.com/chaossat/tiktak/db"
	"github.com/chaossat/tiktak/middleware"
	"github.com/chaossat/tiktak/service/followlist/pb"
	"github.com/spf13/viper"
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
	
	//make([]*pb.User,len(followlist))
}

func main() {

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
