package oss

import (
	"fmt"
	"testing"

	"github.com/chaossat/tiktak/common"
	"github.com/spf13/viper"
)

func TestGetURL(t *testing.T) {
	// 加载配置
	InitConfig()
	// 加载数据库
	common.InitDB()
	fmt.Println(GetURL("videos/42f661c6dcd3ba7a944d4e1ec096e0ada0fd4d66.mp4"))
}

func InitConfig() {
	//workDir, _ := os.Getwd()
	viper.SetConfigName("config.yml")
	viper.SetConfigType("yml")
	// 这里不需要加文件名
	//viper.AddConfigPath(workDir + "/config.yml")
	viper.AddConfigPath("../config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
