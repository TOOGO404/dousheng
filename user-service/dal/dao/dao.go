package dao

import (
	"datasource/database/model"
	"user-service/dal"
)

func RegisterNewUser(user *model.User) error {
	return dal.DB.Create(user).Error
}

func CheckUserPwd(user *model.User) int64 {
	dal.DB.Where(user).First(&user)
	return user.Id
}

func GetUserInfo(uid int64) model.User {
	var user model.User
	dal.DB.Where("id = ?", uid).First(&user)
	return user
}
