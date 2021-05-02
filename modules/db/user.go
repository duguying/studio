// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/5/18.

package db

import (
	"duguying/studio/g"
	"duguying/studio/modules/dbmodels"

	"github.com/gogather/com"
)

func RegisterUser(username string, password string, email string) (user *dbmodels.User, err error) {
	salt := com.RandString(7)
	passwordEncrypt := com.Md5(password + salt)
	tfaSecret := com.RandString(10)
	user = &dbmodels.User{
		Username:  username,
		Password:  passwordEncrypt,
		Salt:      salt,
		Email:     email,
		TfaSecret: tfaSecret,
	}
	err = g.Db.Table("users").Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUser(username string) (user *dbmodels.User, err error) {
	user = &dbmodels.User{}
	err = g.Db.Table("users").Where("username=?", username).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserById(uid uint) (user *dbmodels.User, err error) {
	user = &dbmodels.User{}
	err = g.Db.Table("users").Where("id=?", uid).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CheckUsername(username string) (valid bool, err error) {
	count := int64(0)
	err = g.Db.Table("users").Where("username=?", username).Count(&count).Error
	if err != nil {
		return false, err
	} else {
		return count <= 0, nil
	}
}
