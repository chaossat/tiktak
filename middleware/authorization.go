package middleware

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

//TODO:jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenHeader := context.PostForm("token")
		if tokenHeader == "" {
			context.JSON(http.StatusOK, gin.H{
				"status_code": 1,
				"status_msg":  "用户token不存在",
			})
			context.Abort()
			return
		}
		//checktoken := strings.Split(tokenHeader, " ")
		//if len(checktoken) != 2 && checktoken[0] != "Bearer" {
		//	context.JSON(http.StatusOK, gin.H{
		//		"status_code": 1,
		//		"status_msg":  "token格式错误",
		//	})
		//	context.Abort()
		//	return
		//}
		key, err := CheckToken(tokenHeader)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"status_code": 1,
				"status_msg":  "token错误",
			})
			context.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			context.JSON(http.StatusOK, gin.H{
				"status_code": 1,
				"status_msg":  "token过期",
			})
			context.Abort()
			return
		}
		//context.Set("user_id", key.UserID)
		context.Next()
	}
}
