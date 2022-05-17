package model

func GetVideoList(latest_time int64) []*VideoInfo {
	var videos = []*Video{}
	err := Db.Where("update_time < ?", latest_time).
		Order("update_time desc").
		Limit(30).Find(&videos).Error
	if err != nil {
		return []*VideoInfo{}
	}
	//TODO
	return nil
}
