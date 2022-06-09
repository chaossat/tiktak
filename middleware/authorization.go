package middleware

import (
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

//token验证

//载荷
type MyClaims struct {
	UserID int64 `json:"userID"`
	jwt.StandardClaims
}

var jwtkey []byte

func getkey() []byte {
	jwtkey = []byte(viper.GetString("secretkey"))
	return jwtkey
}

//生成一个token
func CreateToken(ID int64) (string, error) {
	userID := ID
	//到期时间
	expireTime := time.Now().Add(24 * 30 * time.Hour)
	SetClaims := MyClaims{
		UserID: userID,
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

//验证token,并返回解码信息
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
