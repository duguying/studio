// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/1/25.

package model

import "encoding/json"

type User struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	Email    string `json:"email"`
	Verified int    `json:"verified"`
}

func (u *User) String() string {
	user := *u
	user.Password = "*****"
	user.Salt = "*****"

	c, _ := json.Marshal(&user)
	return string(c)
}
