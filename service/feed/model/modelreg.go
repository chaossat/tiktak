package model

import (
	"log"
	"strconv"

	"github.com/go-redis/redis"
)

//获取视频作者信息
func GetUser(userid int64) (User, error) {
	var user User
	err := Db.Where("id = ?", userid).First(&user).Error
	if err != nil {
		log.Println("查找user错误")
		return user, err
	}
	return user, err
}

//获取视频列表
func GetVideoList(latest_time int64) ([]Video, error) {
	var videos = []Video{}
	err := Db.Where("update_time < ?", latest_time).
		Order("update_time desc").
		Limit(30).Find(&videos).Error
	if err != nil {
		if err.Error() == "record not found" {
			return videos, nil
		}
		log.Println("查询视频列表错误")
		return videos, err
	}
	return videos, nil
}

//是否已点赞
func IsFavorite(userid, videoid int64) (bool, error) {
	res, err := Rdb.GetBit(strconv.FormatInt(videoid, 10), (userid - 1)).Result()
	if err != nil {
		log.Println("查询是否已点赞错误", err.Error())
		return false, err
	}

	if res == 1 {
		return true, nil
	}

	return false, nil
}

//获取视频点赞数
func GetFavoriteCount(videoid int64) (int64, error) {
	count := redis.BitCount{Start: 0, End: -1}
	cnt, err := Rdb.BitCount(strconv.FormatInt(videoid, 10), &count).Result()
	if err != nil {
		log.Println("获取视频点赞数错误", err.Error())
		return 0, err
	}
	return cnt, nil
}

//获取视频评论数
func GetCommentCount(video Video) int64 {
	//var video Video
	//Db.Where("id = ?", videoid).First(&video)
	cnt := Db.Model(&video).Association("Comments").Count()
	return int64(cnt)
}

////判断当前用户是否已点赞该视频
//func IsFavorite(user User, video Video) (bool, error) {
//	//var user User
//	//var video Video
//	//Db.Debug().Where("id = ?", userid).First(&user)
//	//Db.Debug().Where("id = ?", videoid).First(&video)
//	err := Db.Debug().Model(&user).Association("FavoriteVideos").Find(&video).Error
//	if err != nil {
//		log.Println(err.Error())
//		return false, err
//	}
//	return true, nil
//}

//用户部分
//获取关注数
func GetFollowCount(user User) int64 {
	var follows int64
	//var user User
	//err := Db.Where("id=?", userid).First(&user).Error
	//if err != nil {
	//	log.Println("查询用户错误")
	//	return -1, err
	//}
	cnt := Db.Model(&user).Association("Follows").Count()
	follows = int64(cnt)
	return follows
}

//获取粉丝数
func GetFollowerCount(user User) int64 {
	var followers int64
	//var user User
	//err := Db.Where("id=?", userid).First(&user).Error
	//if err != nil {
	//	log.Println("查询用户错误")
	//	return -1, err
	//}
	cnt := Db.Model(&user).Association("Followers").Count()
	followers = int64(cnt)
	return followers
}

//判断是否已关注
func IsFollow(user, author User) (bool, error) {
	//var user, author User
	//Db.Where("id = ?", userid).First(&user)
	//Db.Where("id = ?", authorid).First(&author)
	err := Db.Model(&user).Association("Follows").Find(&author).Error
	if err != nil {
		if err.Error() == "record not found" {
			return false, nil
		}
		log.Println("查询是否已关注错误")
		return false, err
	}
	return true, nil
}
