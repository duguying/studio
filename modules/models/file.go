// Copyright 2017. All rights reserved.
// This file is part of blog project
// Created by duguying on 2017/11/2.

package models

import "time"

type File struct {
	Id       int
	Filename string
	Path     string
	Time     time.Time
	Store    string
	Mime     string
}
