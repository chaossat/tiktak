package main

import (
	"os"

	_ "github.com/CodyGuo/godaemon"
	"github.com/chaossat/tiktak/common"
	"github.com/chaossat/tiktak/controller"
	"github.com/chaossat/tiktak/oss"
	"github.com/chaossat/tiktak/router"
	feedmodel "github.com/chaossat/tiktak/service/feed/model"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	r := gin.Default()
	InitConfig()
	router.Init(r)
	common.InitDB()
	feedmodel.InitRedis()
	go oss.Init()

	go controller.DebugInit()

	port := viper.GetString("server.port")
	panic(r.Run(port))
}

//InitConfig:初始化配置文件设置
func InitConfig() {
	workDir, _ := os.Getwd() //获取当前工作路径，非文件路径，以终端显示路径为准
	viper.SetConfigName("config.yml")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
