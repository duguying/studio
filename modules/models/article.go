// Copyright 2017. All rights reserved.
// This file is part of blog project
// Created by duguying on 2017/11/2.

package models

import (
	"encoding/json"
	"strings"
	"time"
)

const (
	ART_STATUS_DRAFT   = 0
	ART_STATUS_PUBLISH = 1
)

type Article struct {
	Id       uint      `json:"id"`
	Title    string    `json:"title"`
	Uri      string    `json:"uri"`
	Keywords string    `json:"keywords"`
	Abstract string    `json:"abstract"`
	Content  string    `json:"content"`
	Author   string    `json:"author"`
	AuthorId uint      `json:"author_id"`
	Time     time.Time `json:"time"`
	Count    uint      `json:"count"`
	Status   int       `json:"status"`
}

func (a *Article) String() string {
	c, _ := json.Marshal(a)
	return string(c)
}

func (a *Article) ToArticleContent() *WrapperArticleContent {
	return &WrapperArticleContent{
		Id:        a.Id,
		Title:     a.Title,
		Uri:       a.Uri,
		Author:    a.Author,
		Tags:      strings.Split(strings.Replace(a.Keywords, "，", ",", -1), ","),
		CreatedAt: a.Time,
		ViewCount: a.Count,
		Content:   a.Content,
	}
}

func (a *Article) ToArticleTitle() *WrapperArticleTitle {
	return &WrapperArticleTitle{
		Id:        a.Id,
		Title:     a.Title,
		Uri:       a.Uri,
		Author:    a.Author,
		CreatedAt: a.Time,
		ViewCount: a.Count,
	}
}
