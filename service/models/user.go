// Copyright 2019. All rights reserved.
// This file is part of duguying project
// I am coding in Tencent
// Created by rainesli on 2019/8/27.

// Package models 接口模型
package models

import "time"

type UserInfo struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar"`
	Access    string    `json:"access"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginArgs struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterArgs struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type ChangePasswordRequest struct {
	Username    string `json:"username"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
