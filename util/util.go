package util

import (
	"crypto/sha1"
	"encoding/hex"
)

//Sha1:计算并返回文件的哈希值
func Sha1(file []byte) string {
	sha1 := sha1.New()
	sha1.Write(file)
	return hex.EncodeToString(sha1.Sum([]byte{}))
}
