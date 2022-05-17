package db

//FavorateCountByID:根据用户id获取关注的数量
func FavorateCountByID(uid int) (int, error) {
	return 0, nil
}

//IsFavorate:根据用户id和视频id查询是否用户点赞了该视频
func IsFavorate(uid, vid int) (bool, error) {
	return false, nil
}
