// Copyright 2017. All rights reserved.
// This file is part of blog project
// Created by duguying on 2017/11/2.

package models

import (
	"encoding/json"
	"time"
)

const (
	ART_STATUS_DRAFT   = 0
	ART_STATUS_PUBLISH = 1
)

type Article struct {
	Id       int64     `json:"id"`
	Title    string    `json:"title"`
	Uri      string    `json:"uri"`
	Keywords string    `json:"keywords"`
	Abstract string    `json:"abstract"`
	Content  string    `json:"content"`
	Author   string    `json:"author"`
	Time     time.Time `json:"time"`
	Count    int       `json:"count"`
	Status   int       `json:"status"`
}

func (a *Article) String() string {
	c, _ := json.Marshal(a)
	return string(c)
}
