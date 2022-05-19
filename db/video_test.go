package db

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/chaossat/tiktak/common"
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
	ves, _ := VideoedByID(1)
	for i := 0; i < len(ves); i++ {
		// 查询
		log.Println(ves[i])
	}

}

func TestVideoByID(t *testing.T) {
	InitConfig()
	// 加载数据库
	common.InitDB()
	res, err := VideoedByID(1)
	fmt.Println("错误信息:", err)
	for _, j := range res {
		fmt.Println(*j)
	}
}

func TestGetVideos(t *testing.T) {
	InitConfig()
	// 加载数据库
	common.InitDB()
	res, err := GetVideos(time.Now().Unix())
	if err != nil {
		fmt.Println("111")
		return
	}
	fmt.Println("获取到的视频:")
	for _, j := range res {
		fmt.Println(*j)
	}
}

func TestVideoCountByID(t *testing.T) {
	InitConfig()
	// 加载数据库
	common.InitDB()
	res, err := VideoCountByID(1)
	fmt.Println("错误信息:", err)
	fmt.Println("获取到的视频数量:", res)
}
