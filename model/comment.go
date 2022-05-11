package model

//Comment:视频评论
type Comment struct {
	ID      int    `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	Content string `gorm:"not null"`
	UserID  int    `gorm:"not null"` //评论者的id
	VideoID int    `gorm:"not null"` //被评论视频的id
}
