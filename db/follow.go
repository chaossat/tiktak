package db

//FollowCountByID:根据用户id获取关注的数量
func FollowCountByID(uid int) (int, error) {
	// user := &model.User_info{}
	// err := common.GetDB().Preload("Follows").Where("id = ?", uid).Find(user).Error
	// return len(user.Follows), err
	return 0, nil
}

//FollowerCountByID:根据用户id获取关注者的数量
func FollowerCountByID(uid int) (int, error) {
	// user := &model.User_info{}
	// err := common.GetDB().Preload("Followers").Where("id = ?", uid).Find(user).Error
	// return len(user.Followers), err
	return 0, nil
}
