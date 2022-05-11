package controller

import (
	"fmt"

	"github.com/chaossat/tiktak/db"
	"github.com/gin-gonic/gin"
)

//LoginHandler:登录接口
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
		LoginResponse(ctx, 0, "Login Failed!", 0, "userVerifyInfo.Token")
		return
	}
	err, userinfo := db.UserInfoByName(username)
	if err != nil {
		fmt.Println(err.Error())
		LoginResponse(ctx, -2, "Error Occoured!", 0, "")
		return
	}
	// userVerifyInfo := model.User_verify_Info{}
	// db.GetUserVerifyInfo(username, &userVerifyInfo)
	// if userVerifyInfo.ID < 1 {
	// 	LoginResponse(ctx, -1, "No Such User!", 0, "")
	// 	return
	// }
	// if util.Encodeing(password) != userVerifyInfo.Password {
	// 	LoginResponse(ctx, -2, "Wrong Password!", 0, "")
	// 	return
	// }
	// userVerifyInfo.ExpirationTime = int(time.Now().Unix()) + 60*60*24 //token过期时间为1天
	// userVerifyInfo.Token = util.TokenGenerator(username)              //构建新的token
	// err := db.UserLogin(&userVerifyInfo)
	err, token := db.CreateToken(&userinfo)
	if err != nil {
		fmt.Println(err.Error())
		LoginResponse(ctx, -3, "Error Occoured!", 0, "")
		return
	}
	LoginResponse(ctx, 0, "Login Succeed!", userinfo.ID, token)
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
