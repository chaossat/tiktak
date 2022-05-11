package db

// import (
// "github.com/chaossat/tiktak/common"
// "github.com/chaossat/tiktak/model"
// )

// //NewUser: 创建新用户
// //TODO:补全功能
// func NewUser(user model.User_info) error {
// 	// return common.GetDB().Create(user).Error
// 	return nil
// }

// //GetUserVerifyInfo:根据用户名获取用户验证信息
// func GetUserVerifyInfo(username string, userVerifyInfo *model.User_verify_Info) error {
// 	return common.GetDB().Where("username = ?", username).First(userVerifyInfo).Error
// }

// //GetUserVerifyInfoWithToken:根据token获取用户验证信息
// func GetUserVerifyInfoWithToken(userVerifyInfo *model.User_verify_Info) error {
// 	return common.GetDB().Where("token = ?", userVerifyInfo.Token).First(userVerifyInfo).Error
// }
