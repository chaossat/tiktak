package db

import (
	"github.com/chaossat/tiktak/common"
	"log"
	"testing"
)

func TestVideoedByID(t *testing.T) {
	InitConfig()
	// 加载配置
	InitConfig()
	// 加载数据库
	common.InitDB()
	// 插入
	//VideoUpload(&model.Video{
	//	Title:         "测试视频3",
	//	AuthorID:      1,
	//	UpdateTime:    time.Now().Second(),
	//	Location:      "d:",
	//	CommentCount:  0,
	//	FavoriteCount: 0,
	//	Play_url:      "",
	//	Cover_url:     "",
	//})
	ves,_:=VideoedByID(1)
	for i := 0; i < len(ves); i++ {
		// 查询
		log.Println(ves[i])
	}

}
