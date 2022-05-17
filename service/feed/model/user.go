package model

//User_info:用户信息
type User struct {
	ID             int64 `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	Name           string
	Follow_count   int64
	Follower_count int64
	Password_Hash  string // 加密后的密码
}

//server端返回的用户信息
type UserInfo struct {
	ID             int64 `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	Name           string
	Follow_count   int64
	Follower_count int64
	IsFollow       bool
}
