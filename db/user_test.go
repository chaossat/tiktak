package db

import (
	"github.com/chaossat/tiktak/common"
	"github.com/chaossat/tiktak/model"
	"github.com/spf13/viper"
	"log"
	"testing"
)

func TestUser_Upload(t *testing.T) {
	// 加载配置
	InitConfig()
	// 加载数据库
	common.InitDB()
	// 上传数据
	err:= UserInfoRegister("xiaom","wajjxuexianajiaq")
	if err!=nil{
		log.Println(err)
	}
}

func TestUser_Update(t *testing.T){
	// 定义结构体
	ds:=&model.User_info{
		ID: 2,
		Name:             "测试二号更新",
		Follow_count:     0,
		Follower_count:   0,
	}
	// 加载配置
	InitConfig()
	// 加载数据库
	common.InitDB()
	// 更新数据
	err:=UserInfoUpdate(ds)
	if err!=nil{
		log.Println(err)
	}
}

func TestUserInfoByName(t *testing.T) {
	// 加载配置
	InitConfig()
	// 加载数据库
	common.InitDB()
	// 查询用户信息根据名称
	err,inf:=UserInfoByName("xiaom")
	if err!=nil{
		log.Println(err)
	}else{
		log.Println(inf)
	}
}

func TestUserInfoById(t *testing.T) {
	// 加载配置
	InitConfig()
	// 加载数据库
	common.InitDB()
	// 查询用户信息根据名称
	err,inf:=UserInfoById(2)
	if err!=nil{
		log.Println(err)
	}else{
		log.Println(inf)
	}
}
//func TestName(t *testing.T) {
//	// 加载配置
//	InitConfig()
//	// 加载数据库
//	common.InitDB()
//	migration.SetAutoMigrate(common.GetDB())
//}


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