package db

import (
	"github.com/chaossat/tiktak/common"
	"github.com/chaossat/tiktak/model"
)

//NewUser: 创建新用户
//TODO:补全功能
func NewUser(user model.User_info) error {
	// return common.GetDB().Create(user).Error
	return nil
}

//GetUserVerifyInfo:根据用户名获取用户验证信息
func GetUserVerifyInfo(username string, userVerifyInfo *model.User_verify_Info) error {
	return common.GetDB().Where("username = ?", username).First(userVerifyInfo).Error
}

//GetUserVerifyInfoWithToken:根据token获取用户验证信息
func GetUserVerifyInfoWithToken(userVerifyInfo *model.User_verify_Info) error {
	return common.GetDB().Where("token = ?", userVerifyInfo.Token).First(userVerifyInfo).Error
}

//UserLogin:用户登录成功，更新用户持有的token和过期时间
func UserLogin(userVerifyInfo *model.User_verify_Info) error {
	err := common.GetDB().Model(userVerifyInfo).Where("id = ?", userVerifyInfo.ID).Update("token", userVerifyInfo.Token).Error
	if err != nil {
		return err
	}
	err = common.GetDB().Model(userVerifyInfo).Where("id = ?", userVerifyInfo.ID).Update("expirationtime", userVerifyInfo.ExpirationTime).Error
	return err
}
