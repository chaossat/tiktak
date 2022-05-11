package common

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"testing"
)

type ID struct {
	uid int
}

func (receiver ID) TableName() string {
	return "test"
}

func TestDB(t *testing.T) {
	InitConfig()
	id:=ID{}
	fmt.Println("加载配置完成")
	db:=InitDB()
	fmt.Println(db==nil)
	db.Debug().Select("uid").First(&id)
	log.Println(id)
}

func TestRedis(t*testing.T){
	InitConfig()
	err:=InitRedis()
	if err!=nil{
		log.Println(err.Error())
	}
	log.Println(rdb.Get("f1"))
}

func InitConfig() {
	viper.SetConfigName("config.yml")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

