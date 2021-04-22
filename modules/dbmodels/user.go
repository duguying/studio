// Copyright 2017. All rights reserved.
// This file is part of blog project
// Created by duguying on 2017/11/2.

package dbmodels

import (
	"duguying/studio/g"
	"duguying/studio/service/models"
	"time"

	"github.com/gogather/json"
)

const (
	RoleUser  = 0
	RoleAdmin = 1
)

var role = []string{"user", "admin"}

type User struct {
	Id        uint      `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Salt      string    `json:"salt"`
	Email     string    `json:"email"`
	Role      int       `json:"role" gorm:"default:0"`
	TfaSecret string    `json:"tfa_secret"` // 2FA secret base 32
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) String() string {
	c, _ := json.Marshal(u)
	return string(c)
}

func (u *User) ToInfo() *models.UserInfo {
	host := g.Config.Get("system", "host", "http://duguying.net")
	return &models.UserInfo{
		Id:       u.Id,
		Username: u.Username,
		Email:    u.Email,
		Avatar:   host + "/logo.png",
		Access:   role[u.Role],
	}
}
