package util

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net"
	"os/exec"
	"strings"
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

// MD5:对密码加盐后进行md5加密
func MD5V(pwd string) string {
	return MD5(pwd + salt)
}

//GetIP:获取本机ip地址
func GetIP() (ip string) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}

//CoverGenerator:为视频截取第一帧作为封面，需要安装ffmpeg环境
func CoverGenerator(videoPath string, coverName string) {
	exec.Command("ffmpeg", "-i", videoPath, "tempimage/"+coverName).Run()
}
