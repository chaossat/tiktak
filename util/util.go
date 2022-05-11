package util

import (
	"crypto/sha1"
	"encoding/hex"
	"strconv"
	"time"
)

var salt = "6824Nil"

//Sha1:计算并返回文件的哈希值
func Sha1(file []byte) string {
	sha1 := sha1.New()
	sha1.Write(file)
	return hex.EncodeToString(sha1.Sum([]byte{}))
}

//Encoding:加盐后计算哈希值
func Encodeing(s string) string {
	return Sha1([]byte(s + salt))
}

//TokenGenerator:根据用户名生成token
func TokenGenerator(username string) string {
	return Encodeing(username + strconv.Itoa(int(time.Now().Unix())) + salt)
}
