package db

import (
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

//VideoedByID：根据用户id找到这个用户已发布的视频列表
func VideoedByID(uid int) ([]*model.Video, error) {
	var videoes = []*model.Video{}
	err := common.GetDB().Where("author_id=?", uid).Find(&videoes).Error
	return videoes, err
}

//VideoCountByID:根据用户id获取发布的视频数量
func VideoCountByID(uid int) (int, error) {
	// user := &model.User_info{}
	// err := common.GetDB().Preload("Videos").Where("id = ?", uid).Find(user).Error
	// return len(user.Videos), err
	return 0, nil
}
