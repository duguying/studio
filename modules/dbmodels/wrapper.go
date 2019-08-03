// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/1/25.

package dbmodels

import (
	"github.com/gogather/json"
	"time"
)

type WrapperArticleContent struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Uri       string    `json:"uri"`
	Author    string    `json:"author"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	ViewCount uint      `json:"view_count"`
	Content   string    `json:"content"`
}

func (ac *WrapperArticleContent) String() string {
	c, _ := json.Marshal(ac)
	return string(c)
}

type WrapperArticleTitle struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Uri       string    `json:"uri"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	ViewCount uint      `json:"view_count"`
}

func (at *WrapperArticleTitle) String() string {
	c, _ := json.Marshal(at)
	return string(c)
}
