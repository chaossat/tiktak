package db

import (
	"errors"

	"github.com/chaossat/tiktak/common"
	"github.com/chaossat/tiktak/model"
)

//CommentCountByVID:根据视频id获取评论数量
func CommentCountByVID(vid int) (int, error) {
	video := &model.Video{}
	common.GetDB().Where("id = ?", vid).First(video)
	if video.ID == 0 {
		return -1, errors.New("no such video")
	}
	cnt := common.GetDB().Model(video).Association("comments").Count()
	return cnt, nil
}

//CommentsByVID:根据视频id获取评论列表，根据接口要求，按照时间倒序直接返回所有评论
func CommentsByVID(vid int) ([]*model.Comment, error) {
	comments := []*model.Comment{}
	err := common.GetDB().Where("video_id = ?", vid).Order("update_time desc").Find(comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}
