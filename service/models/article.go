// Copyright 2019. All rights reserved.
// This file is part of duguying project
// I am coding in Tencent
// Created by rainesli on 2019/8/29.

package models

import "github.com/gogather/json"

type Article struct {
	Id       uint     `json:"id"`
	Title    string   `json:"title"`
	Uri      string   `json:"uri"`
	Keywords []string `json:"keywords"`
	Abstract string   `json:"abstract"`
	Content  string   `json:"content"`
	Type     int      `json:"type"`
	Draft    bool     `json:"draft"`
}

func (aar *Article) String() string {
	c, _ := json.Marshal(aar)
	return string(c)
}
