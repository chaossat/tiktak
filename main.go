package main

import (
	"os"

	"github.com/chaossat/tiktak/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//TODO:在router.go中设置路由规则
//TODO:补全代码
func main() {
	r := gin.Default()
	InitConfig()
	router.Init(r)
	// port := viper.GetString("server.port")
	// panic(r.Run())
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
