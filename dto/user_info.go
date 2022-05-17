package dto

import (
	"github.com/chaossat/tiktak/model"
)

//TODO:完成其他符合接口形式的结构体
type User_info_dto struct {
	code int    `json:"status_code"`
	msg  string `json:"status_msg"`
	user User   `json:"user"`
}

type User struct {
	id             int    `json:"id"`
	name           string `json:"name"`
	follow_count   int    `json:"follow_count"`
	follower_count int    `json:"follower_count"`
	is_follow      bool   `json:"is_follow"`
}

// 根据数据库中的user_info，构建并返回User_info_dto
func Dto_user_info(user *model.User_info) User_info_dto {
	user_inf := User{
		id:             int(user.ID),
		name:           user.Name,
		follow_count:   00000000000000, //待修改
		follower_count: 00000000000000, //待修改
		is_follow:      false,
	}
	return User_info_dto{
		code: 200,
		msg:  "success",
		user: user_inf,
	}
}
