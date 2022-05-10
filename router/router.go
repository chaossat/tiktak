package router

import (
	"net/http"

	"github.com/chaossat/tiktak/controller"
	"github.com/gin-gonic/gin"
)

//TODO:设置路由规则
func Init(r *gin.Engine) {
	r.StaticFS("tempvideo", http.Dir("./tempfile")) //设置静态资源
	//设置路由规则
	douyin := r.Group("/douyin")
	{
		publish := douyin.Group("/publish")
		{
			publish.POST("action", controller.UploadHandler)
		}
	}
}
