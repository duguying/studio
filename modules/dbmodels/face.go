// Copyright 2020. All rights reserved.
// This file is part of duguying project
// I am coding in Tencent
// Created by rainesli on 2020/5/8.

package dbmodels

import "time"

type Face struct {
	Id             uint      `json:"id"`
	FileId         uint      `json:"file_id"`
	FaceDescriptor string    `json:"face_descriptor" sql:"type:longtext"`
	LabelId        uint      `json:"label_id"`
	CreatedAt      time.Time `json:"created_at"`
}
