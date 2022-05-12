package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"log"
	"time"
)

//token验证

//载荷
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtkey string

func getkey() string {
	jwtkey = viper.GetString("secretkey")
	return jwtkey
}

//生成一个token
func CreateToken(username string) (string, error) {
	//到期时间
	expireTime := time.Now().Add(24 * 30 * time.Hour)
	SetClaims := MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "douyin",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaim.SignedString([]byte(getkey()))
	if err != nil {
		log.Println("token err:", err.Error())
		return "", err
	}
	return token, nil
}

//验证token
func CheckToken(token string) (*MyClaims, error) {
	setToken, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return getkey(), nil
	})
	if err != nil {
		return nil, err
	}
	if key, _ := setToken.Claims.(*MyClaims); setToken.Valid {
		return key, nil
	}
	return nil, errors.New("token验证失败")
}

//TODO:jwt中间件
