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

type ArticleContentListResponse struct {
	Ok    bool              `json:"ok"`
	Msg   string            `json:"msg"`
	Total uint              `json:"total"`
	List  []*ArticleContent `json:"list"`
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

type ArticleContentGetResponse struct {
	Ok   bool            `json:"ok"`
	Msg  string          `json:"msg"`
	Data *ArticleContent `json:"data"`
}

