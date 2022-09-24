// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/6/30.

package db

import (
	"duguying/studio/modules/dbmodels"
	"duguying/studio/utils"
	"path"
	"time"

	"gorm.io/gorm"
)

func SaveFile(tx *gorm.DB, fpath string, mime string, size uint64, md5 string, userID uint, fileType dbmodels.FileType) (f *dbmodels.File, err error) {
	filename := path.Base(fpath)
	f = &dbmodels.File{
		Filename:  filename,
		Path:      fpath,
		Store:     dbmodels.LOCAL,
		Mime:      mime,
		Size:      size,
		FileType:  fileType,
		Md5:       md5,
		UserID:    userID,
		CreatedAt: time.Now(),
	}
	err = tx.Model(dbmodels.File{}).Create(f).Error
	if err != nil {
		return nil, err
	} else {
		return f, nil
	}
}

// DeleteFile 删除文件
func DeleteFile(tx *gorm.DB, id string) (err error) {
	return tx.Model(dbmodels.File{}).Where("id=?", id).Delete(&dbmodels.File{}).Error
}

// CheckFileRef 检查文件引用
func CheckFileRef(tx *gorm.DB, file *dbmodels.File) (cnt int64, err error) {
	url := utils.GetFileURL(file.Path)
	cnt, err = FileCountArticleRef(tx, url)
	if err != nil {
		return 0, err
	}
	coverRefCnt, err := FileCountCoverRef(tx, file.ID)
	if err != nil {
		return 0, err
	}
	return cnt + coverRefCnt, nil
}

// FileCountCoverRef 封面文件引用计数
func FileCountCoverRef(tx *gorm.DB, fileID string) (cnt int64, err error) {
	err = tx.Model(dbmodels.Cover{}).Where("file_id=?", fileID).Count(&cnt).Error
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

// PageFile 文件分页列表
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

// ListAllMediaFile 列举媒体文件
func ListAllMediaFile(tx *gorm.DB, userID uint) (list []*dbmodels.File, err error) {
	list = []*dbmodels.File{}
	where := "file_type in (?)"
	params := []interface{}{[]int{int(dbmodels.FileTypeImage), int(dbmodels.FileTypeVideo)}}
	if userID > 0 {
		where = where + " and user_id=?"
		params = append(params, userID)
	}
	err = tx.Model(dbmodels.File{}).Where(where, params...).Order("created_at desc").Find(&list).Error
	if err != nil {
		return nil, err
	} else {
		return list, nil
	}
}

// UpdateFileMediaSize 更新媒体文件尺寸
func UpdateFileMediaSize(tx *gorm.DB, fileID string, width, height int) (err error) {
	return tx.Model(dbmodels.File{}).Where("id=?", fileID).Updates(map[string]interface{}{
		"media_width":  width,
		"media_height": height,
	}).Error
}

// UpdateFileThumbneil 更新媒体缩略图
func UpdateFileThumbneil(tx *gorm.DB, fileID string, thumbneil string) (err error) {
	return tx.Model(dbmodels.File{}).Where("id=?", fileID).Updates(map[string]interface{}{
		"thumbnail": thumbneil,
	}).Error
}
