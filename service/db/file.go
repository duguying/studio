// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/6/30.

package db

import (
	"duguying/studio/g"
	"duguying/studio/modules/models"
	"path"
	"time"
)

func SaveFile(fpath string, mime string, size uint64) (err error) {
	filename := path.Base(fpath)
	f := &models.File{
		Filename:  filename,
		Path:      fpath,
		Store:     models.LOCAL,
		Mime:      mime,
		Size:      size,
		CreatedAt: time.Now(),
	}
	errs := g.Db.Table("files").Create(f).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return errs[0]
	} else {
		return nil
	}
}

func PageFile(page uint64, size uint64) (list []*models.File, total uint, err error) {
	list = []*models.File{}
	total = 0
	errs := g.Db.Table("files").Count(&total).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return nil, 0, errs[0]
	}

	errs = g.Db.Table("files").Order("id desc").Offset((page - 1) * size).Limit(size).Find(&list).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return nil, 0, errs[0]
	} else {
		return list, total, nil
	}
}
