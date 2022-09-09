// Copyright 2017. All rights reserved.
// This file is part of blog project
// Created by duguying on 2017/11/2.

package dbmodels

import (
	"database/sql/driver"
	"duguying/studio/service/models"
	"duguying/studio/utils"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/gogather/json"
)

const (
	LOCAL StorageType = 0
	OSS   StorageType = 1

	FileTypeUnknown FileType = 0
	FileTypeImage   FileType = 1
	FileTypeVideo   FileType = 2
	FileTypeArchive FileType = 3

	RecognizeNotNeed RecognizeStatus = 0
	RecognizeDone    RecognizeStatus = 1
)

var (
	FileTypeMap = map[FileType]string{
		FileTypeUnknown: "unknown",
		FileTypeImage:   "image",
		FileTypeVideo:   "video",
		FileTypeArchive: "archive",
	}
)

type StorageType int64

func (tt StorageType) Value() (driver.Value, error) {
	return int64(tt), nil
}

func (tt *StorageType) Scan(value interface{}) error {
	val, ok := value.(int64)
	if !ok {
		switch reflect.TypeOf(value).String() {
		case "[]uint8":
			{
				ba := []byte{}
				for _, b := range value.([]uint8) {
					ba = append(ba, byte(b))
				}
				val, _ = strconv.ParseInt(string(ba), 10, 64)
				break
			}
		default:
			{
				return fmt.Errorf("value: %v, is not int, is %s", value, reflect.TypeOf(value))
			}
		}
	}
	*tt = StorageType(val)
	return nil
}

// ---------

type FileType int64

func (tt FileType) Value() (driver.Value, error) {
	return int64(tt), nil
}

func (tt *FileType) Scan(value interface{}) error {
	val, ok := value.(int64)
	if !ok {
		switch reflect.TypeOf(value).String() {
		case "[]uint8":
			{
				ba := []byte{}
				for _, b := range value.([]uint8) {
					ba = append(ba, byte(b))
				}
				val, _ = strconv.ParseInt(string(ba), 10, 64)
				break
			}
		default:
			{
				return fmt.Errorf("value: %v, is not int, is %s", value, reflect.TypeOf(value))
			}
		}
	}
	*tt = FileType(val)
	return nil
}

// ---------

type RecognizeStatus int64

func (rs RecognizeStatus) Value() (driver.Value, error) {
	return int64(rs), nil
}

func (rs *RecognizeStatus) Scan(value interface{}) error {
	val, ok := value.(int64)
	if !ok {
		switch reflect.TypeOf(value).String() {
		case "[]uint8":
			{
				ba := []byte{}
				for _, b := range value.([]uint8) {
					ba = append(ba, byte(b))
				}
				val, _ = strconv.ParseInt(string(ba), 10, 64)
				break
			}
		default:
			{
				return fmt.Errorf("value: %v, is not int, is %s", value, reflect.TypeOf(value))
			}
		}
	}
	*rs = RecognizeStatus(val)
	return nil
}

// ---------

type File struct {
	UUID

	Filename   string          `json:"filename"`
	Path       string          `json:"path"`
	Store      StorageType     `json:"store"`
	Mime       string          `json:"mime"`
	Size       uint64          `json:"size"`
	FileType   FileType        `json:"file_type" gorm:"default:0" sql:"comment:'文件类型'"`
	Md5        string          `json:"md5" sql:"comment:'MD5'"`
	Recognized RecognizeStatus `json:"recognized" gorm:"default:0" sql:"comment:'识别状态'"`
	UserID     uint            `json:"user_id" gorm:"comment:'文件所有者';index"`
	CreatedAt  time.Time       `json:"created_at"`
}

func (f *File) String() string {
	c, _ := json.Marshal(f)
	return string(c)
}

func (f *File) ToModel() *models.File {
	return &models.File{
		ID:         f.ID,
		Filename:   f.Filename,
		Path:       f.Path,
		Store:      int64(f.Store),
		Mime:       f.Mime,
		Size:       f.Size,
		FileType:   int64(f.FileType),
		Md5:        f.Md5,
		Recognized: int64(f.Recognized),
		UserID:     f.UserID,
		CreatedAt:  f.CreatedAt,
	}
}

func (f *File) ToMediaFile() *models.MediaFile {
	return &models.MediaFile{
		ID:        f.ID,
		Filename:  f.Filename,
		URL:       utils.GetFileURL(f.Path),
		Mime:      f.Mime,
		Size:      f.Size,
		FileType:  FileTypeMap[f.FileType],
		Md5:       f.Md5,
		UserID:    f.UserID,
		CreatedAt: f.CreatedAt,
	}
}
