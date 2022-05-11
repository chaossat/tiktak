package model

//Favorate：点赞
type Favorate struct {
	ID   int `gorm:"column:id;primary_key;not null"` //点赞id
	From int `gorm:"not null"`                       //点赞的用户id
	To   int `gorm:"not null"`                       //被点赞的视频id
}
