package model

//Video:视频信息
type Video struct {
	ID             int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	Title          string
	AuthorID       int64  //作者id
	FavorateCounts int64  //点赞数
	FavorateList   string //点赞列表
	CommentCounts  int64  //评论数
	CommentList    string //评论列表
	UpdateTime     int64  //发布、更新时间
	Location       string //储存位置
}
