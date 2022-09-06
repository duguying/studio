package models

import "time"

type File struct {
	ID         string    `json:"id"`
	Filename   string    `json:"filename"`
	Path       string    `json:"path"`
	Store      int64     `json:"store"`
	Mime       string    `json:"mime"`
	Size       uint64    `json:"size"`
	FileType   int64     `json:"file_type"`
	Md5        string    `json:"md5"`
	Recognized int64     `json:"recognized"`
	LocalExist bool      `json:"local_exist"`
	UserID     uint      `json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
}
