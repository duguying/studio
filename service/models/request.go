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
