package router

import (
	"net/http"

	"github.com/chaossat/tiktak/controller"
	"github.com/gin-gonic/gin"
)

//TODO:设置路由规则
func Init(r *gin.Engine) {
	r.StaticFS("tempfile", http.Dir("./tempfile"))   //设置静态资源
	r.StaticFS("tempvideo", http.Dir("./tempvideo")) //设置静态资源
	r.StaticFS("static", http.Dir("./static"))       //设置静态资源
	//设置分组路由规则
	douyin := r.Group("/douyin")
	{
		douyin.GET("/feed", controller.Feed)
		douyin.GET("/user/", controller.UserInfHandler)
		user := douyin.Group("/user")
		{
			user.POST("/register/", controller.Register)
			user.POST("/login/", controller.LoginHandler)
		}
		publish := douyin.Group("/publish")
		{
			publish.POST("/action/", controller.UploadHandler)
			publish.GET("/list", controller.PubListHandler)
		}
		favorite := douyin.Group("/favorite")
		{
			favorite.GET("/list/", controller.FavoriteList)
			favorite.POST("/action/", controller.FavoriteAction)
		}
		relation := douyin.Group("/relation")
		{
			relation.GET("/follower/list/", controller.Followerlist)
			relation.GET("/follow/list/", controller.FollowListHandler)
			relation.POST("/action/", controller.FollowActionHandler)
		}
		comment := douyin.Group("/comment")
		{
			comment.POST("/action/", controller.CommentActionHandler)
			comment.GET("/list/", controller.CommentListHandler)
		}
	}
}
