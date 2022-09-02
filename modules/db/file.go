// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/6/30.

package db

import (
	"duguying/studio/modules/dbmodels"
	"path"
	"time"

	"gorm.io/gorm"
)

func SaveFile(tx *gorm.DB, fpath string, mime string, size uint64, md5 string, userID uint) (err error) {
	filename := path.Base(fpath)
	f := &dbmodels.File{
		Filename:  filename,
		Path:      fpath,
		Store:     dbmodels.LOCAL,
		Mime:      mime,
		Size:      size,
		Md5:       md5,
		UserID:    userID,
		CreatedAt: time.Now(),
	}
	err = tx.Table("files").Create(f).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

func PageFile(tx *gorm.DB, page uint64, size uint64) (list []*dbmodels.File, total int64, err error) {
	list = []*dbmodels.File{}
	total = 0
	err = tx.Table("files").Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = tx.Table("files").Order("created_at desc").Offset(int((page - 1) * size)).Limit(int(size)).Find(&list).Error
	if err != nil {
		return nil, 0, err
	} else {
		return list, total, nil
	}
}
