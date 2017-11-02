// Copyright 2017. All rights reserved.
// This file is part of blog project
// Created by duguying on 2017/11/2.

package models

import "time"

type Varify struct {
	Id       int
	Username string
	Code     string
	Overdue  time.Time
}
