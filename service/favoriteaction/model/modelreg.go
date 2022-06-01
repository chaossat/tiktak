package model

import (
	"strconv"
)

//点赞
func GiveLike(userid int64, videoid int64) (bool, error) {
	res, err := Rdb.GetBit(strconv.FormatInt(videoid, 10), (userid - 1)).Result()
	if err != nil {
		return false, err
	}

	if res == 1 {
		return true, nil
	}

	res, err = Rdb.SetBit(strconv.FormatInt(videoid, 10), (userid - 1), 1).Result()
	if err != nil {
		return false, err
	}
	var user User
	var video Video
	Db.Where("id = ?", userid).First(&user)
	Db.Where("id = ?", videoid).First(&video)
	err = Db.Model(&user).
		Association("FavoriteVideos").
		Append(&video).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

//取消点赞
func CancelLike(userid int64, videoid int64) (bool, error) {
	res, err := Rdb.GetBit(strconv.FormatInt(videoid, 10), (userid - 1)).Result()
	if err != nil {
		return false, err
	}

	if res == 0 {
		return true, nil
	}

	res, err = Rdb.SetBit(strconv.FormatInt(videoid, 10), (userid - 1), 0).Result()
	if err != nil {
		return false, err
	}

	var user User
	var video Video
	Db.Where("id = ?", userid).First(&user)
	Db.Where("id = ?", videoid).First(&video)
	err = Db.Model(&user).
		Association("FavoriteVideos").
		Append(&video).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
