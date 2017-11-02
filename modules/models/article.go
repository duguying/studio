// Copyright 2017. All rights reserved.
// This file is part of blog project
// Created by duguying on 2017/11/2.

package models

import "time"

type Article struct {
	Id       int
	Title    string
	Uri      string
	Keywords string
	Abstract string
	Content  string
	Author   string
	Time     time.Time
	Count    int
	Status   int
}

const (
	ART_STATUS_DRAFT   = 0
	ART_STATUS_PUBLISH = 1
)
