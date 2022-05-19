package db

import (
	"time"
)

// token失效时间
const TOKEN_EXPIRETIME = time.Hour * 10

//
///**
//给用户颁发token：token为key，用户信息为value，存储在redis中
//*/
//func CreateToken(userinf *model.User_info) (error, string) {
//	// to do: 是否考虑在mysql持久化token
//	key := util.Encodeing(uuid.NewV4().String())
//	value, err_json := json.Marshal(userinf)
//	if err_json != nil {
//		return err_json, ""
//	}
//	err_set := common.GetRDB().Set(key, value, TOKEN_EXPIRETIME).Err()
//	if err_set != nil {
//		return err_set, ""
//	}
//	return err_set, key
//}
