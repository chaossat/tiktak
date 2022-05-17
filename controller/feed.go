package controller

import (
	"github.com/chaossat/tiktak/middleware"
	"github.com/gin-gonic/gin"
	"log"
	"reflect"
)

func Feed(ctx *gin.Context) {
	latest_time_str := ctx.DefaultQuery("latest_time", "current")
	token := ctx.Query("token")
	_, err := middleware.CheckToken(token)
	if err != nil {
		
	}
	log.Println(reflect.TypeOf(latest_time_str))
	//if latest_time_str == "current" {
	//	latest_time_str = time.Now().String()
	//}
	//latest_time, err := strconv.ParseInt(latest_time_str, 10, 64)
	//if err != nil {
	//	ctx.JSON(200, gin.H{
	//		"status_code": 1,
	//		"status_msg":  "时间错误",
	//	})
	//	return
	//}
}
