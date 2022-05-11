package controller

import (
	"time"

	"github.com/chaossat/tiktak/db"
	"github.com/chaossat/tiktak/model"
	"github.com/chaossat/tiktak/util"
	"github.com/gin-gonic/gin"
)

//LoginHandler:登录接口
func LoginHandler(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	userVerifyInfo := model.User_verify_Info{}
	db.GetUserVerifyInfo(username, &userVerifyInfo)
	if userVerifyInfo.ID < 1 {
		LoginResponse(ctx, -1, "No Such User!", 0, "")
		return
	}
	if util.Encodeing(password) != userVerifyInfo.Password {
		LoginResponse(ctx, -2, "Wrong Password!", 0, "")
		return
	}
	userVerifyInfo.ExpirationTime = int(time.Now().Unix()) + 60*60*24 //token过期时间为1天
	userVerifyInfo.Token = util.TokenGenerator(username)              //构建新的token
	err := db.UserLogin(&userVerifyInfo)
	if err != nil {
		LoginResponse(ctx, -3, "Error Occoured While Updating User Token!", 0, "")
		return
	}
	LoginResponse(ctx, 0, "Login Succeed!", userVerifyInfo.ID, userVerifyInfo.Token)
}

//LoginResponse:返回登录处理信息
func LoginResponse(ctx *gin.Context, code int, msg string, user_id int, token string) {
	ctx.JSON(200, gin.H{
		"status_code": code,
		"status_msg":  msg,
		"user_id":     user_id,
		"token":       token,
	})
}
