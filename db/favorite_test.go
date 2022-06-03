package db

import (
	"fmt"
	"testing"

	"github.com/chaossat/tiktak/common"
)

func TestFavoriteCountByID(t *testing.T) {
	// 加载配置
	InitConfig()
	// 加载数据库
	common.InitDB()
	fmt.Println("查询结果：")
	fmt.Println(FavoriteCountByID(5))
}
