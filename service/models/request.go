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

// type CommonGetterRequest struct {
// 	Id uint `json:"id" form:"id"`
// }

type MonthlyPagerRequest struct {
	Page  uint `json:"page" form:"page"`
	Size  uint `json:"size" form:"size"`
	Year  uint `json:"year" form:"year"`
	Month uint `json:"month" form:"month"`
}

type TopGetterRequest struct {
	Top uint `json:"top" form:"top"`
}

type ArticleUriGetterRequest struct {
	Uri string `json:"uri" form:"uri"`
	Id  uint   `json:"id" form:"id"`
}

type ArticlePublishRequest struct {
	Id      uint `json:"id" form:"id"`
	Publish bool `json:"publish" form:"publish"`
}

type TagPagerRequest struct {
	Page uint   `json:"page" form:"page"`
	Size uint   `json:"size" form:"size"`
	Tag  string `json:"tag" form:"tag"`
}

type SearchPagerRequest struct {
	Page    uint   `json:"page" form:"page"`
	Size    uint   `json:"size" form:"size"`
	Keyword string `json:"keyword" form:"keyword"`
}

type IntGetter struct {
	ID uint `json:"id" form:"id" binding:"required"`
}

type StringGetter struct {
	ID string `json:"id" form:"id" binding:"required"`
}

type UserIDGetter struct {
	UserID uint `json:"user_id" form:"user_id" binding:"required"`
}

type FileSyncRequest struct {
	FileID  string `json:"file_id" form:"file_id"`
	CosType string `json:"cos_type" form:"cos_type"`
}
