package model

//User_info:用户信息
type User_info struct {
	ID             int    `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	Name           string `gorm:"not null"`
	Follow_count   int
	Follower_count int
	Favorates      string //点赞的视频
	Comments       string //发布的评论
}

//User_verify_Info：用户验证信息
type User_verify_Info struct {
	ID             int    `gorm:"column:id;primary_key;not null"`
	Name           string `gorm:"not null"`
	Password       string `gorm:"not null"`
	Token          string //用户当前应持有的token
	ExpirationTime int    //token的过期时间
}
