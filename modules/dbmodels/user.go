// Copyright 2017. All rights reserved.
// This file is part of blog project
// Created by duguying on 2017/11/2.

package dbmodels

import (
	"duguying/studio/g"
	"duguying/studio/service/models"
	"duguying/studio/utils"
	"time"

	"github.com/gogather/json"
)

const (
	RoleUser  = 0
	RoleAdmin = 1
)

var role = []string{"user", "admin"}

type User struct {
	ID            uint      `json:"id"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	Salt          string    `json:"salt"`
	Email         string    `json:"email"`
	Role          int       `json:"role" gorm:"default:0"`
	TfaSecret     string    `json:"tfa_secret"` // 2FA secret base 32
	AvatarFileID  string    `json:"vatar_file_id" gorm:"comment:'图像文件ID';index"`
	AvatarFileKey string    `json:"avatar_file_key" gorm:"comment:'图像文件路径'"`
	CreatedAt     time.Time `json:"created_at"`
}

func (u *User) String() string {
	c, _ := json.Marshal(u)
	return string(c)
}

func (u *User) ToInfo() *models.UserInfo {
	host := g.Config.Get("system", "host", "http://duguying.net")
	avatar := host + "/logo.png"
	if u.AvatarFileKey != "" {
		avatar = utils.GetFileURL(u.AvatarFileKey)
	}
	return &models.UserInfo{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Avatar:   avatar,
		Access:   role[u.Role],
	}
}
