package db

import (
	"github.com/chaossat/tiktak/common"
	"github.com/chaossat/tiktak/model"
)

// 更新用户信息
func UserInfoUpdate(info *model.User_info)error{
	return common.GetDB().Model(info).Where("id=?",info.ID).Update(info).Error
}
// 加入用户信息
func UserInfoUpload(info *model.User_info) error {
	return common.GetDB().Create(info).Error
}
// 查询用户信息根据id
func UserInfoById(id int)(error,model.User_info){
	inf:=model.User_info{}
	err:=common.GetDB().Where("id=?",id).First(&inf).Error
	return err,inf
}
// 查询用户信息根据用户名
func UserInfoByName(name string)(error,model.User_info)  {
	inf:=model.User_info{}
	err:=common.GetDB().Where("name=?",name).First(&inf).Error
	return err,inf
}
