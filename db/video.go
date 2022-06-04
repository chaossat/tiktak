package db

import (
	"errors"

	"github.com/chaossat/tiktak/common"
	"github.com/chaossat/tiktak/model"
)

// VideoUpload: 插入视频信息
func VideoUpload(video *model.Video) error {
	return common.GetDB().Create(video).Error
}

//VideoLocationUpdate: 更新视频的地址信息
func VideoLocationUpdate(video *model.Video) error {
	return common.GetDB().Model(video).Where("id = ?", video.ID).Update("Location", video.Location).Error
}

//CoverLocationUpdate:更新封面的地址信息
func CoverLocationUpdate(video *model.Video) error {
	return common.GetDB().Model(video).Where("id = ?", video.ID).Update("Cover_location", video.Cover_location).Error
}

//VideoedByID：根据用户id找到这个用户已发布的视频列表,根据接口要求，直接按时间倒序列出所有视频
func VideoedByID(uid int) ([]*model.Video, error) {
	var videos = []*model.Video{}
	err := common.GetDB().Where("author_id=?", uid).Order("update_time desc").Find(&videos).Error
	return videos, err
}

//GetVideos:根据输入的时间戳，返回在该时间戳之前的 30个最新视频
func GetVideos(timeStamp int64) ([]*model.Video, error) {
	var videos []*model.Video
	err := common.GetDB().Where("update_time < ?", timeStamp).Order("update_time desc").Limit(30).Find(&videos).Error
	return videos, err
}

//VideoCountByID:根据用户id获取发布的视频数量
func VideoCountByID(uid int) (int, error) {
	user := &model.User{}
	common.GetDB().Where("id = ?", uid).First(user)
	if user.ID == 0 {
		return -1, errors.New("no such user")
	}
	cnt := common.GetDB().Model(user).Association("Videos").Count()
	return cnt, nil
}

//IsVideoExist查询对应id的视频是否存在
func IsVideoExist(vid int) (bool, error) {
	p := 0
	err := common.GetDB().Model(&model.Video{}).Where("id = ?", vid).Count(&p).Error
	if err != nil {
		return false, err
	}
	return p == 1, nil
}

//删除视频，仅调试用
func DeleteVideoVyID(vid int) error {
	return common.GetDB().Where("id = ?", vid).Delete(&model.Video{}).Error
}
