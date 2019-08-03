// Copyright 2017. All rights reserved.
// This file is part of blog project
// Created by duguying on 2017/11/2.

package dbmodels

import (
	"github.com/gogather/json"
	"strings"
	"time"
)

const (
	ArtStatus_Draft   = 0
	ArtStatus_Publish = 1
	ArtStatus_Delete  = 2
)

const (
	ContentType_HTML     = 0
	ContentType_MarkDown = 1
)

type Article struct {
	Id          uint      `json:"id"`
	Title       string    `json:"title"`
	Uri         string    `json:"uri"`
	Keywords    string    `json:"keywords"`
	Abstract    string    `json:"abstract"`
	Type        int       `json:"type"`
	Content     string    `json:"content"`
	Author      string    `json:"author"`
	AuthorId    uint      `json:"author_id"`
	Count       uint      `json:"count"`
	Status      int       `json:"status"`
	PublishTime time.Time `json:"publish_time"`
	CreatedAt   time.Time `json:"created_at"`
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
		CreatedAt: a.CreatedAt,
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
		CreatedAt: a.CreatedAt,
		ViewCount: a.Count,
	}
}
