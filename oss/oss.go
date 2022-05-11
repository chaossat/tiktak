package oss

import (
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/chaossat/tiktak/db"
	"github.com/chaossat/tiktak/model"
	"github.com/spf13/viper"
)

//通信结构体
type VideoOBJ struct {
	OssPath   string
	File      *os.File
	VideoMeta model.Video
}

var ossCli *oss.Client
var ossBucket *oss.Bucket

//MQ_channel:用于传送储存请求
//不使用集群，所以不使用MQ
var MQ_channel = make(chan *VideoOBJ, 100)

//获取oss的Client
func GetClient() *oss.Client {
	if ossCli != nil {
		return ossCli
	}
	ossCli, err := oss.New(viper.GetString("datasource.oss.OSSEndpoint"), viper.GetString("datasource.oss.OSSAccesskeyID"), viper.GetString("datasource.oss.OSSAccesskeySecret"))
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return ossCli
}

//获取oss上Client的bucket存储空间
func Bucket() *oss.Bucket {
	if ossBucket != nil {
		return ossBucket
	}
	cli := GetClient()
	if cli == nil {
		return nil
	}
	bucket, err := cli.Bucket(viper.GetString("datasource.oss.OSSBucket"))
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	ossBucket = bucket
	return ossBucket
}

//Init:持续消费MQ_channel里的转存请求
func Init() {
	for {
		video := <-MQ_channel
		video.File.Seek(0, 0) // 游标重新回到File文件头部，否则oss读不出任何数据
		err := Bucket().PutObject(video.OssPath, video.File)
		video.File.Close()
		if err != nil {
			fmt.Printf("Failed while pushing to oss, err:%s\n", err.Error())
			continue
		}
		err = db.VideoLocationUpdate(&(video.VideoMeta))
		if err != nil {
			fmt.Printf("Failed to update mysql, err:%s\n", err.Error())
			continue
		}
		err = os.Remove("./tempfile/" + video.OssPath[7:])
		if err != nil {
			fmt.Printf("Failed to delete local file, err:%s\n", err.Error())
		}
	}
}
