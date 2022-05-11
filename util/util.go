package util

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

//Sha1:计算并返回文件的哈希值
func Sha1(file []byte) string {
	sha1 := sha1.New()
	sha1.Write(file)
	return hex.EncodeToString(sha1.Sum([]byte{}))
}
// MD5:对密码进行md5加密
func MD5V(pwd string) string  {
	// 创建切片
	data := []byte(pwd)
	// 获取md5的数组
	has := md5.Sum(data)
	// 转化为16进制
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

