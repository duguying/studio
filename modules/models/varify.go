// Copyright 2017. All rights reserved.
// This file is part of blog project
// Created by duguying on 2017/11/2.

package models

import (
	"encoding/json"
	"time"
)

type Varify struct {
	Id       int       `json:"id"`
	Username string    `json:"username"`
	Code     string    `json:"code"`
	Overdue  time.Time `json:"overdue"`
}

func (v *Varify) String() string {
	c, _ := json.Marshal(v)
	return string(c)
}
