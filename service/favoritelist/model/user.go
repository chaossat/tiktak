package model

//User_info:用户信息
type User struct {
	ID            int64 `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	Username      string
	Password_Hash string // 加密后的密码
	//用户关注列表，多对多，gorm自动生成一个中间表
	Follows []*User `gorm:"many2many:user_follows;association_jointable_foreignkey:follow_id"` //该用户的关注
	//粉丝列表，多对多，同样自动生成一个中间表
	Followers []*User `gorm:"many2many:user_followers;association_jointable_foreignkey:follower_id"` //该用户的粉丝
	//用户发布的视频，通过视频表的作者id作为外键关联
	Videos []Video `gorm:"foreignkey:AuthorID"` //该用户发布的作品
	//用户点赞过的视频，多对多，生成一个中间表，
	FavoriteVideos []Video `gorm:"many2many:user_favorite;association_jointable_foreignkey:favorite_id"` //用户点赞过的作品
}
