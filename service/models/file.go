package models

import "time"

type File struct {
	ID              string    `json:"id"`
	Filename        string    `json:"filename"`
	Path            string    `json:"path"`
	Store           int64     `json:"store"`
	Mime            string    `json:"mime"`
	Size            uint64    `json:"size"`
	FileType        int64     `json:"file_type"`
	Md5             string    `json:"md5"`
	Recognized      int64     `json:"recognized"`
	LocalExist      bool      `json:"local_exist"`
	ArticleRefCount int       `json:"article_ref_count"`
	COS             bool      `json:"cos"`
	UserID          uint      `json:"user_id"`
	CreatedAt       time.Time `json:"created_at"`
}

type MediaFile struct {
	ID        string    `json:"id"`
	Filename  string    `json:"filename"`
	URL       string    `json:"url"`
	Mime      string    `json:"mime"`
	Size      uint64    `json:"size"`
	FileType  string    `json:"file_type" `
	Md5       string    `json:"md5" sql:"comment:'MD5'"`
	UserID    uint      `json:"user_id" gorm:"comment:'文件所有者';index"`
	CreatedAt time.Time `json:"created_at"`
}
