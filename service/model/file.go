// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/1/25.

package model

import (
	"encoding/json"
	"time"
)

type File struct {
	Id       uint      `json:"id"`
	Filename string    `json:"filename"`
	Path     string    `json:"path"`
	Time     time.Time `json:"time"`
	Store    int       `json:"store"`
	Mime     string    `json:"mime"`
}

func (f *File) String() string {
	c, _ := json.Marshal(f)
	return string(c)
}
