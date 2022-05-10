package model

//User_info:用户信息
type User_info struct {
	ID               int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	Name             string
	Follow_count     int
	Follower_count   int
	Favorates        string //点赞的视频
	Videos_published string //发布的视频
	Comments         string //发布的评论
}
