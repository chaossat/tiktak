package main

import (
	"context"
	"fmt"

	"github.com/chaossat/tiktak/db"
	"github.com/chaossat/tiktak/middleware"
	"github.com/chaossat/tiktak/service/login/pb"
)

type LoginHandler struct {
	pb.UnimplementedLoginServer
}

//Login:grpc的Login接口
func (loginHandler *LoginHandler) Login(ctx context.Context, loginRequest *pb.DouyinUserLoginRequest) (*pb.DouyinUserLoginResponse, error) {
	username := loginRequest.GetUsername()
	password := loginRequest.GetPassword()
	ok, err := db.UserLogin(username, password)
	if err != nil {
		fmt.Println(err.Error())
		return LoginResponse(-1, "Error Occoured!", 0, ""), nil
	}
	if !ok {
		return LoginResponse(-2, "Invalid Verify Imformation!", 0, ""), nil
	}
	err, userinfo := db.UserInfoByName(username)
	if err != nil {
		fmt.Println(err.Error())
		return LoginResponse(-3, "Error Occoured!", 0, ""), nil
	}
	token, err := middleware.CreateToken(int(userinfo.ID))
	if err != nil {
		fmt.Println(err.Error())
		return LoginResponse(-4, "Error Occoured!", 0, ""), nil
	}
	return LoginResponse(0, "Login Succeed!", int(userinfo.ID), token), nil
}

//LoginResponse:返回登录处理信息
func LoginResponse(code int, msg string, user_id int, token string) *pb.DouyinUserLoginResponse {
	codeResponse := int32(code)
	userID := int64(user_id)
	return &pb.DouyinUserLoginResponse{
		StatusCode: &codeResponse,
		StatusMsg:  &msg,
		UserId:     &userID,
		Token:      &token,
	}
}
