// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/5/18.

package db

import (
	"duguying/studio/modules/dbmodels"

	"github.com/gogather/com"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func RegisterUser(tx *gorm.DB, username string, password string, email string) (user *dbmodels.User, err error) {
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
	err = tx.Table("users").Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UserChangePassword 修改密码
func UserChangePassword(tx *gorm.DB, username string, newPassword string) (err error) {
	newSalt := com.Md5(uuid.New().String())
	newPasswd := com.Md5(newPassword + newSalt)

	return tx.Model(dbmodels.User{}).Where("username=?", username).Updates(map[string]interface{}{
		"salt":     newSalt,
		"password": newPasswd,
	}).Error
}

func GetUser(tx *gorm.DB, username string) (user *dbmodels.User, err error) {
	user = &dbmodels.User{}
	err = tx.Table("users").Where("username=?", username).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByID(tx *gorm.DB, uid uint) (user *dbmodels.User, err error) {
	user = &dbmodels.User{}
	err = tx.Table("users").Where("id=?", uid).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CheckUsername(tx *gorm.DB, username string) (valid bool, err error) {
	count := int64(0)
	err = tx.Table("users").Where("username=?", username).Count(&count).Error
	if err != nil {
		return false, err
	} else {
		return count <= 0, nil
	}
}
