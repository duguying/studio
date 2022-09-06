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

// DeleteFile 删除文件
func DeleteFile(tx *gorm.DB, id string) (err error) {
	return tx.Model(dbmodels.File{}).Where("id=?", id).Delete(&dbmodels.File{}).Error
}

func PageFile(tx *gorm.DB, page uint64, size uint64, userID uint) (list []*dbmodels.File, total int64, err error) {
	list = []*dbmodels.File{}
	total = 0
	err = tx.Model(dbmodels.File{}).Where("user_id=?", userID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = tx.Model(dbmodels.File{}).Where("user_id=?", userID).Order("created_at desc").Offset(int((page - 1) * size)).Limit(int(size)).Find(&list).Error
	if err != nil {
		return nil, 0, err
	} else {
		return list, total, nil
	}
}

// GetFile 获取文件信息
func GetFile(tx *gorm.DB, id string) (file *dbmodels.File, err error) {
	file = &dbmodels.File{}
	err = tx.Model(dbmodels.File{}).Where("id=?", id).First(file).Error
	if err != nil {
		return nil, err
	}
	return file, nil
}
