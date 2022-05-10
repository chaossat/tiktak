package oss

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
)

var ossCli *oss.Client
var ossBucket *oss.Bucket

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
