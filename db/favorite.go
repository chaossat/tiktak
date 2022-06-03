package db

import (
	"errors"

	"github.com/chaossat/tiktak/common"
	"github.com/chaossat/tiktak/model"
)

//FavorateCountByID:根据用户id获取点赞的数量
func FavoriteCountByID(uid int) (int, error) {
	user := &model.User{}
	common.GetDB().Where("id = ?", uid).First(user)
	if user.ID == 0 {
		return -1, errors.New("no such user")
	}
	cnt := common.GetDB().Model(user).Association("FavoriteVideos").Count()
	return cnt, nil
}
