package model

//Comment:视频评论
type Comment struct {
	ID         int64  `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	Content    string `gorm:"not null"`
	UserID     int64  `gorm:"not null"` //评论者的id
	VideoID    int64  `gorm:"not null"` //被评论视频的id
	UpdateTime int64  `gorm:"not null"` //评论的发布、更新时间
}
