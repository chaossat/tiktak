package util

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

var salt = "6824Nil"

//Sha1:计算并返回文件的哈希值
func Sha1(file []byte) string {
	sha1 := sha1.New()
	sha1.Write(file)
	return hex.EncodeToString(sha1.Sum([]byte{}))
}

//MD5:返回对字符串的MD5计算值
func MD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

//Encoding:加盐后计算哈希值
func Encodeing(s string) string {
	return Sha1([]byte(s + salt))
}

// //TokenGenerator:根据用户名生成token
// func TokenGenerator(username string) string {
// 	return Encodeing(username + strconv.Itoa(int(time.Now().Unix())) + salt)
// }

// MD5:对密码加盐后进行md5加密
func MD5V(pwd string) string {
	return MD5(pwd + salt)
}
