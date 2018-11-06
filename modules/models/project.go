// Copyright 2017. All rights reserved.
// This file is part of blog project
// Created by duguying on 2017/11/2.

package models

import (
	"github.com/gogather/json"
	"time"
)

type Project struct {
	Id          uint     `json:"id"`
	Name        string    `json:"name"`
	IconUrl     string    `json:"icon_url"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	Time        time.Time `json:"time"`
}

func (p *Project) String() string {
	c, _ := json.Marshal(p)
	return string(c)
}
