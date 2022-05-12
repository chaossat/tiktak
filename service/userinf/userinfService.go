package userinf

import (
	"context"
	"github.com/chaossat/tiktak/db"
	"github.com/chaossat/tiktak/middleware"
	"github.com/chaossat/tiktak/service/userinf/pb"
)

type UserInf struct {
	pb.UnimplementedUserInfServer
}
/**
获取用户信息服务
 */
func (this*UserInf) GetUserinf(context context.Context, request *pb.DouyinUserRequest) ( *pb.DouyinUserResponse,error) {
	uid:=request.UserId
	token:=request.Token
	_,err:=middleware.CheckToken(*token)
	if err!=nil{
		var code int32=-1
		var msg string="token认证失败"
		response:=pb.DouyinUserResponse{
			StatusCode: &code,
			StatusMsg:  &msg,
			User:       nil,
		}
		return &response,nil
	}
	// 查询数据库获取用户信息
	err,inf:=db.UserInfoById(int(*uid))
	if err!=nil{
		var code int32=-2
		var msg string="找不到用户id"
		response:=pb.DouyinUserResponse{
			StatusCode: &code,
			StatusMsg:  &msg,
			User:       nil,
		}
		return &response,nil
	}
	// 用户信息结构体
	var isFollow=false
	var follow_count int64=int64(inf.Follow_count)
	var follower_count int64=int64(inf.Follower_count)
	var user=pb.User{
		Id:            uid,
		Name:          &inf.Name,
		FollowCount:   &follow_count,
		FollowerCount: &follower_count,
		IsFollow:      &isFollow,
	}
	var code int32=0
	var msg string="成功！"
	return &pb.DouyinUserResponse{
		StatusCode: &code,
		StatusMsg:  &msg,
		User:       &user,
	},nil
}


