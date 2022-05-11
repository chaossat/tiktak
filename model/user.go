package model

//User_info:用户信息
type User_info struct {
	ID             int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	Name           string
	Follow_count   int
	Follower_count int
	Password_Hash  string // 加密后的密码
}
