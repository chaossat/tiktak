package model

import (
	"errors"
)

func CheckUser(username string) bool {
	var user User_info
	Db.Where("username=?", username).First(&user)
	//if err == nil {
	//	log.Println(err.Error(), "校验用户出错")
	//	return true
	//}
	if user.ID > 0 {
		return true
	}
	return false
}

func SaveUser(username, password_hash string) (*User_info, error) {
	if CheckUser(username) {
		return nil, errors.New("用户名已存在")
	}
	user := User_info{
		Username:      username,
		Password_hash: password_hash,
	}
	return &user, Db.Create(&user).Error
}
