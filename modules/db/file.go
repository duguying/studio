// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/6/30.

package db

import (
	"duguying/studio/g"
	"duguying/studio/modules/dbmodels"
	"path"
	"time"
)

func SaveFile(fpath string, mime string, size uint64, md5 string) (err error) {
	filename := path.Base(fpath)
	f := &dbmodels.File{
		Filename:  filename,
		Path:      fpath,
		Store:     dbmodels.LOCAL,
		Mime:      mime,
		Size:      size,
		Md5:       md5,
		CreatedAt: time.Now(),
	}
	err = g.Db.Table("files").Create(f).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

func PageFile(page uint64, size uint64) (list []*dbmodels.File, total int64, err error) {
	list = []*dbmodels.File{}
	total = 0
	err = g.Db.Table("files").Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = g.Db.Table("files").Order("id desc").Offset(int((page - 1) * size)).Limit(int(size)).Find(&list).Error
	if err != nil {
		return nil, 0, err
	} else {
		return list, total, nil
	}
}
