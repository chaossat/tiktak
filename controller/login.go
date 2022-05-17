package controller

import (
	"fmt"

	"github.com/chaossat/tiktak/db"
	"github.com/chaossat/tiktak/middleware"
	"github.com/gin-gonic/gin"
)

//LoginHandler:登录接口处理器
func LoginHandler(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	ok, err := db.UserLogin(username, password)
	if err != nil {
		fmt.Println(err.Error())
		LoginResponse(ctx, -1, "Error Occoured During Verification!", 0, "")
		return
	}
	if !ok {
		LoginResponse(ctx, -2, "Login Failed!", 0, "")
		return
	}
	err, userinfo := db.UserInfoByName(username)
	if err != nil {
		fmt.Println(err.Error())
		LoginResponse(ctx, -3, "Error Occoured!", 0, "")
		return
	}
	token, err := middleware.CreateToken(userinfo.ID)
	if err != nil {
		fmt.Println(err.Error())
		LoginResponse(ctx, -4, "Error Occoured!", 0, "")
		return
	}
	LoginResponse(ctx, 0, "Login Succeed!", int(userinfo.ID), token)
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
