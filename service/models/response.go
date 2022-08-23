// Copyright 2019. All rights reserved.
// This file is part of duguying project
// I am coding in Tencent
// Created by rainesli on 2019/8/29.

package models

type CommonResponse struct {
	Ok  bool   `json:"ok"`
	Msg string `json:"msg"`
}

type CommonCreateResponse struct {
	Ok  bool   `json:"ok"`
	Msg string `json:"msg"`
	Id  uint   `json:"id"`
}

type CommonListResponse struct {
	Ok    bool        `json:"ok"`
	Msg   string      `json:"msg"`
	Total uint        `json:"total"`
	List  interface{} `json:"list"`
}

type CommonSearchListResponse struct {
	Ok    bool        `json:"ok"`
	Msg   string      `json:"msg"`
	Total uint        `json:"total"`
	List  interface{} `json:"list"`
}

type ArticleContentListResponse struct {
	Ok    bool                  `json:"ok"`
	Msg   string                `json:"msg"`
	Total uint                  `json:"total"`
	List  []*ArticleShowContent `json:"list"`
}

type ArticleTitleListResponse struct {
	Ok    bool            `json:"ok"`
	Msg   string          `json:"msg"`
	Total uint            `json:"total"`
	List  []*ArticleTitle `json:"list"`
}

type ArticleArchListResponse struct {
	Ok   bool        `json:"ok"`
	Msg  string      `json:"msg"`
	List []*ArchInfo `json:"list"`
}

type ArticleShowContentGetResponse struct {
	Ok   bool                `json:"ok"`
	Msg  string              `json:"msg"`
	Data *ArticleShowContent `json:"data"`
}

type ArticleContentGetResponse struct {
	Ok   bool            `json:"ok"`
	Msg  string          `json:"msg"`
	Data *ArticleContent `json:"data"`
}

type UserInfoResponse struct {
	Ok   bool      `json:"ok"`
	Msg  string    `json:"msg"`
	Data *UserInfo `json:"data"`
}

type LoginResponse struct {
	Ok  bool   `json:"ok"`
	Msg string `json:"msg"`
	Sid string `json:"sid"`
}

type UploadResponse struct {
	Ok   bool   `json:"ok"`
	Msg  string `json:"msg"`
	Path string `json:"path"`
}
