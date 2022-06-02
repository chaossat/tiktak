package model

//Video:视频信息
type Video struct {
	ID             int64  `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	Title          string `gorm:"not null"`
	AuthorID       int64  `gorm:"not null;index:index_desc"` //作者id
	UpdateTime     int64  `gorm:"not null;index:index_desc"` //发布、更新时间
	Location       string `gorm:"not null"`                  //储存位置
	Cover_location string `gorm:"not null"`                  //封面储存位置
	// PlayLocation   string //视频播放地址
	//视频评论，通过评论表的视频id作为外键关联
	Comments []Comment `gorm:"foreignkey:VideoID"` //视频的评论
}
