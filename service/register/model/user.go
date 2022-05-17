package model

//User_info:用户信息
type User_info struct {
	ID               int64  `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	Username         string `json:"username"`
	Password_hash    string `json:"password_Hash"`
	Follow_count     int64
	Follower_count   int64
	Favorates        []*Video //点赞的视频
	Videos_published []*Video //发布的视频
	//Comments         string   //发布的评论
	//TODO:用户之间的关注关系
}
