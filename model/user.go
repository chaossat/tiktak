package model

//User_info:用户信息
type User_info struct {
	ID   int64 `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	Name string
	//Follow_count   int64
	//Follower_count int64
	Password_Hash  string      // 加密后的密码
	Follows        []User_info //该用户的关注
	Followers      []User_info //该用户的粉丝
	Videos         []Video     //该用户发布的作品
	FavoriteVideos []Video     //用户点赞过的作品
}
