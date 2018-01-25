// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/1/25.

package model

import (
	"encoding/json"
	"time"
)

type Article struct {
	Id       uint      `json:"id"`
	Title    string    `json:"title"`
	Uri      string    `json:"uri"`
	Keywords string    `json:"keywords"`
	Abstract string    `json:"abstract"`
	Content  string    `json:"content"`
	AuthorId uint      `json:"author_id"`
	Time     time.Time `json:"time"`
	Count    uint      `json:"count"`
	Status   int       `json:"status"`
}

func (a *Article) String() string {
	c, _ := json.Marshal(a)
	return string(c)
}
