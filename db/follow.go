package db

import (
	"errors"
	"fmt"

	"github.com/chaossat/tiktak/common"
	"github.com/chaossat/tiktak/model"
)

//FollowCountByID:根据用户id获取关注的数量
func FollowCountByID(uid int) (int64, error) {
	user := &model.User{}
	common.GetDB().Where("id = ?", uid).First(user)
	if user.ID == 0 {
		return -1, errors.New("no such user")
	}
	cnt := common.GetDB().Model(user).Association("Follows").Count()
	return int64(cnt), nil
}

//根据用户的id 爱豆的列表
func FollowListByID(uid int) (userlist []*model.User, err error) {
	user := &model.User{}
	common.GetDB().Where("id = ?", uid).First(user)
	if user.ID == 0 {
		return nil, errors.New("no such user")
	}
	users := []*model.User{}
	err = common.GetDB().Model(user).Association("Follows").Find(&users).Error
	return users, err
}

//FollowerCountByID:根据用户id获取关注者的数量
func FollowerCountByID(uid int) (int64, error) {
	user := &model.User{}
	common.GetDB().Where("id = ?", uid).First(user)
	if user.ID == 0 {
		return -1, errors.New("no such user")
	}
	cnt := common.GetDB().Model(user).Association("Followers").Count()
	return int64(cnt), nil
}

//根据用户的id获取关注者的列表
func FollowerListByID(uid int) (userlist []*model.User, err error) {
	user := &model.User{}
	common.GetDB().Where("id = ?", uid).First(user)
	if user.ID == 0 {
		return nil, errors.New("no such user")
	}
	users := []*model.User{}
	err = common.GetDB().Model(user).Association("Followers").Find(&users).Error
	return users, err
}

//判断是否已关注作者
func IsFollow(user, author model.User) bool {
	err := common.GetDB().Model(&user).Association("Follows").Find(&author).Error
	if err != nil {
		if err.Error() == "record not found" {
			return false
		}
		fmt.Println("查询是否已关注错误", err)
		return false
	}
	return true
}
