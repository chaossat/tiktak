package db

import (
	"fmt"
	"testing"

	"github.com/chaossat/tiktak/common"
)

func TestUserInfo(t *testing.T) {
	// 加载配置
	InitConfig()
	// 加载数据库
	common.InitDB()
	res, err := UserInfo(0)
	fmt.Println(err)
	for _, j := range res {
		fmt.Println(j)
	}
}

func TestInsertUserByID(t *testing.T) {
	// 加载配置
	InitConfig()
	// 加载数据库
	common.InitDB()
	fmt.Println(InsertUserByID(1))
}
