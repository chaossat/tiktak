package db

import (
	"fmt"
	"testing"

	"github.com/chaossat/tiktak/common"
	"github.com/chaossat/tiktak/model"
)

func TestIsFollow(t *testing.T) {
	// 加载配置
	InitConfig()
	// 加载数据库
	common.InitDB()
	// 更新数据
	user1 := model.User{
		ID: 1,
	}
	user2 := model.User{
		ID: 2,
	}
	res := IsFollow(user1, user2)
	fmt.Println("查询结果：", res)
}
