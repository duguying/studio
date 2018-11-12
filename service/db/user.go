// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/5/18.

package db

import (
	"duguying/studio/g"
	"duguying/studio/modules/models"
	"github.com/gogather/com"
)

func RegisterUser(username string, password string, email string) (user *models.User, err error) {
	salt := com.RandString(7)
	passwordEncrypt := com.Md5(password + salt)
	tfaSecret := com.RandString(10)
	user = &models.User{
		Username:  username,
		Password:  passwordEncrypt,
		Salt:      salt,
		Email:     email,
		TfaSecret: tfaSecret,
	}
	errs := g.Db.Table("users").Create(user).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return nil, errs[0]
	}
	return user, nil
}

func GetUser(username string) (user *models.User, err error) {
	user = &models.User{}
	errs := g.Db.Table("users").Where("username=?", username).First(user).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return nil, errs[0]
	}
	return user, nil
}

func GetUserById(uid uint) (user *models.User, err error) {
	user = &models.User{}
	errs := g.Db.Table("users").Where("id=?", uid).First(user).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return nil, errs[0]
	}
	return user, nil
}

func CheckUsername(username string) (valid bool, err error) {
	count := 0
	errs := g.Db.Table("users").Where("username=?", username).Count(&count).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return false, errs[0]
	} else {
		return count <= 0, nil
	}
}
