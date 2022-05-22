package router

import (
	"github.com/chaossat/tiktak/middleware"
	"net/http"

	"github.com/chaossat/tiktak/controller"
	"github.com/gin-gonic/gin"
)

//TODO:设置路由规则
func Init(r *gin.Engine) {
	r.StaticFS("tempfile", http.Dir("./tempfile"))   //设置静态资源
	r.StaticFS("tempvideo", http.Dir("./tempvideo")) //设置静态资源
	//设置分组路由规则
	douyin := r.Group("/douyin")
	{
		douyin.GET("/feed", controller.Feed)
		user := douyin.Group("/user")
		{
			user.POST("/register/", controller.Register)
			user.POST("login", controller.LoginHandler)
		}
		douyin.Use(middleware.JwtToken())
		douyin.GET("/osstest", controller.GetURL) //临时测试地址
		//douyin.GET("/feed", controller.VideoListHandler)
		publish := douyin.Group("/publish")
		{
			publish.POST("action", controller.UploadHandler)
		}
	}
}
