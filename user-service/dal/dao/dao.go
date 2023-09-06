package dao

import (
	"datasource/database/model"
	"errors"
	"fmt"
	"user-service/dal"
	"utils"
)

func RegisterNewUser(user *model.User) error {
	return dal.DB.Create(user).Error
}

func CheckUserPwd(user *model.User) (int64, error) {
	var _user model.User
	err := dal.DB.Where(" email = ? ", user.Email).First(&_user).Error
	if err != nil {
		return 0, err
	}
	fmt.Print(_user.Pwd)
	if utils.ComparePwd(_user.Pwd, user.Pwd) {
		return _user.Id, nil
	} else {
		return 0, errors.New("password error")
	}
}

func GetUserInfo(uid int64) model.User {
	var user model.User
	dal.DB.Where("id = ?", uid).First(&user)
	return user
}
