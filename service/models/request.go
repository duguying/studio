// Copyright 2019. All rights reserved.
// This file is part of duguying project
// I am coding in Tencent
// Created by rainesli on 2019/8/29.

package models

type CommonPagerRequest struct {
	Page   uint `json:"page" form:"page"`
	Size   uint `json:"size" form:"size"`
	Status int  `json:"status" form:"status"`
}

type CommonGetterRequest struct {
	Id uint `json:"id" form:"id"`
}
