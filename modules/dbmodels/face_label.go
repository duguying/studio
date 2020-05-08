// Copyright 2020. All rights reserved.
// This file is part of duguying project
// I am coding in Tencent
// Created by rainesli on 2020/5/8.

package dbmodels

import "time"

type FaceLabel struct {
	Id        uint      `json:"id"`
	Label     string    `json:"label"`
	CreatedAt time.Time `json:"created_at"`
}
