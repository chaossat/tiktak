package db

import (
	"errors"

	"github.com/chaossat/tiktak/common"
	"github.com/chaossat/tiktak/model"
)

//FollowCountByID:根据用户id获取关注的数量
func FollowCountByID(uid int) (int, error) {
	user := &model.User{}
	common.GetDB().Where("id = ?", uid).First(user)
	if user.ID == 0 {
		return -1, errors.New("no such user")
	}
	cnt := common.GetDB().Model(user).Association("Follows").Count()
	return cnt, nil
}

//FollowerCountByID:根据用户id获取关注者的数量
func FollowerCountByID(uid int) (int, error) {
	user := &model.User{}
	common.GetDB().Where("id = ?", uid).First(user)
	if user.ID == 0 {
		return -1, errors.New("no such user")
	}
	cnt := common.GetDB().Model(user).Association("Followers").Count()
	return cnt, nil
}
