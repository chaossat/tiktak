package main

import (
	"os"

	"github.com/chaossat/tiktak/common"
	"github.com/chaossat/tiktak/oss"
	"github.com/chaossat/tiktak/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	r := gin.Default()
	InitConfig()
	router.Init(r)
	common.InitDB()
	go oss.Init()
	port := viper.GetString("server.port")
	panic(r.Run(port))
}

//InitConfig:初始化配置文件设置
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config.yml")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
