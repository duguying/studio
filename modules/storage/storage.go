// Copyright 2019. All rights reserved.
// This file is part of duguying project
// I am coding in Tencent
// Created by rainesli on 2020/4/11.

package storage

type FileInfo struct {
	Path string `json:"path"`
	Size int64  `json:"size"`
	Mime string `json:"mime"`
}

type Storage interface {
	List(remotePrefix string) (list []*FileInfo, err error)
	GetFileInfo(remotePath string) (info *FileInfo, err error)
	AddFile(localPath string, remotePath string) (err error)
	RenameFile(remotePath string, newRemotePath string) (err error)
	RemoveFile(remotePath string) (err error)
	FetchFile(remotePath string, localPath string) (err error)
}
