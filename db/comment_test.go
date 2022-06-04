package db

import (
	"fmt"
	"testing"

	"github.com/chaossat/tiktak/common"
)

func TestCommentCountByVID(t *testing.T) {
	// 加载配置
	InitConfig()
	// 加载数据库
	common.InitDB()
	for i := 9; i < 30; i++ {
		fmt.Println("vid:", i)
		fmt.Println(CommentCountByVID(i))
	}
}
