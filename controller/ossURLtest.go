package controller

import (
	"github.com/chaossat/tiktak/oss"
	"github.com/gin-gonic/gin"
)

//测试oss.GetURL
func GetURL(ctx *gin.Context) {
	//测试oss链接
	URL := oss.GetURL("videos/c90b44a96ed080c1a6c8ce8888a40a5aaaa7e7ca.mp4")
	URL2 := oss.GetURL("images/c90b44a96ed080c1a6c8ce8888a40a5aaaa7e7ca.jpg")
	//测试本地链接
	// URL := oss.GetURL("tempfile/c90b44a96ed080c1a6c8ce8888a40a5aaaa7e7ca.mp4")
	ctx.JSON(200, gin.H{
		"video": URL,
		"cover": URL2,
	})
}
