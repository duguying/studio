// Copyright 2017. All rights reserved.
// This file is part of blog project
// Created by duguying on 2017/11/2.

package models

import (
	"encoding/json"
	"time"
)

type File struct {
	Id        uint      `json:"id"`
	Filename  string    `json:"filename"`
	Path      string    `json:"path"`
	Store     string    `json:"store"`
	Mime      string    `json:"mime"`
	CreatedAt time.Time `json:"created_at"`
}

func (f *File) String() string {
	c, _ := json.Marshal(f)
	return string(c)
}
