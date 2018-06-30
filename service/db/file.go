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
