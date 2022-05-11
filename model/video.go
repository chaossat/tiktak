package model

//Video:视频信息
type Video struct {
	ID         int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	Title      string
	AuthorID   int    //作者id
	UpdateTime int    //发布、更新时间
	Location   string //储存位置
}
