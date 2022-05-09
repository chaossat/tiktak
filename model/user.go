package model

//TODO：完成其他需要储存在mysql的结构体
type User_info struct {
	id               int
	name             string
	follow_count     int
	follower_count   int
	favorates        string //点赞的视频
	videos_published string //发布的视频
	comments         string //发布的评论
}
