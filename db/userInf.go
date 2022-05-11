package db

import (
	"github.com/chaossat/tiktak/common"
	"github.com/chaossat/tiktak/model"
	"github.com/chaossat/tiktak/util"
)

// 更新用户信息
func UserInfoUpdate(info *model.User_info) error {
	return common.GetDB().Model(info).Where("id=?", info.ID).Update(info).Error
}

// 查询用户信息根据id
func UserInfoById(id int) (error, model.User_info) {
	inf := model.User_info{}
	err := common.GetDB().Where("id=?", id).First(&inf).Error
	return err, inf
}

// 查询用户信息根据用户名
func UserInfoByName(name string) (error, model.User_info) {
	inf := model.User_info{}
	err := common.GetDB().Where("name=?", name).First(&inf).Error
	return err, inf
}

// 用户登录
func UserLogin(name string, pwd string) (bool, error) {
	p := 0
	// 密码做MD5
	pwd_md5 := util.MD5V(pwd)
	err := common.GetDB().Where("name=? and password_hash=?", name, pwd_md5).Count(&p).Error
	return p == 1, err
}

// 用户注册
func UserInfoRegister(name string, pwd string) error {
	// 密码做md5
	pwd_md5 := util.MD5V(pwd)
	inf := model.User_info{
		Name:           name,
		Follow_count:   0,
		Follower_count: 0,
		Password_Hash:  pwd_md5,
	}
	err := common.GetDB().Create(&inf).Error
	return err
}
