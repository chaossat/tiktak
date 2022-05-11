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
