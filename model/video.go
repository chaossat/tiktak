package model

//Video:视频信息
type Video struct {
	ID            int64 `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	Title         string
	AuthorID      int    //作者id
	UpdateTime    int    //发布、更新时间
	Location      string //储存位置
	CommentCount  int64  //评论数
	FavoriteCount int64  //点赞数
	Play_url      string //播放地址
	Cover_url	  string //封面地址
}
