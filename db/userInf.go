package db

import (
	"errors"
	"strings"

	"github.com/chaossat/tiktak/common"
	"github.com/chaossat/tiktak/model"
	"github.com/chaossat/tiktak/util"
)

// 更新用户信息
func UserInfoUpdate(info *model.User) error {
	return common.GetDB().Model(info).Where("id=?", info.ID).Update(info).Error
}

// 查询用户信息根据id
func UserInfoById(id int) (model.User, error) {
	inf := model.User{}
	err := common.GetDB().Where("id=?", id).First(&inf).Error
	return inf, err
}

// 查询用户信息根据用户名
func UserInfoByName(name string) (error, model.User) {
	inf := model.User{}
	if strings.Contains(name, " ") {
		return errors.New("invalid username"), inf
	}
	err := common.GetDB().Where("username=?", name).First(&inf).Error
	return err, inf
}

// 用户登录
func UserLogin(name string, pwd string) (bool, error) {
	if strings.Contains(name, " ") || strings.Contains(pwd, " ") {
		return false, errors.New("invalid username or password")
	}
	p := 0
	// 密码做MD5
	pwd_md5 := util.MD5V(pwd)
	err := common.GetDB().Model(&model.User{}).Where("username=? and password_hash=?", name, pwd_md5).Count(&p).Error
	return p == 1, err
}

// 用户注册
func UserInfoRegister(name string, pwd string) error {
	if strings.Contains(name, " ") || strings.Contains(pwd, " ") {
		return errors.New("invalid username or password")
	}
	// 密码做md5
	pwd_md5 := util.MD5V(pwd)
	inf := model.User{
		Username:      name,
		Password_Hash: pwd_md5,
	}
	err := common.GetDB().Create(&inf).Error
	return err
}
