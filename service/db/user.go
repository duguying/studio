// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/5/18.

package db

import (
	"duguying/studio/g"
	"duguying/studio/modules/models"
)

func GetUser(username string) (user *models.User, err error) {
	user = &models.User{}
	errs := g.Db.Table("users").Where("username=?", username).First(user).GetErrors()
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
