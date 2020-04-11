// Copyright 2019. All rights reserved.
// This file is part of duguying project
// I am coding in Tencent
// Created by rainesli on 2020/4/11.

package storage

type QcloudCos struct {
	ak     string
	sk     string
	bucket string
}

func NewQcloudOss(ak string, sk string, bucket string) (storage *QcloudCos, err error) {
	storage = &QcloudCos{
		ak:     ak,
		sk:     sk,
		bucket: bucket,
	}
	return storage, nil
}

func (q QcloudCos) List(remotePrefix string) (list []*FileInfo, err error) {
	panic("implement me")
}

func (q QcloudCos) GetFileInfo(remotePath string) (info *FileInfo, err error) {
	panic("implement me")
}

func (q QcloudCos) AddFile(localPath string, remotePath string) (err error) {
	panic("implement me")
}

func (q QcloudCos) RenameFile(remotePath string, newRemotePath string) (err error) {
	panic("implement me")
}

func (q QcloudCos) RemoveFile(remotePath string) (err error) {
	panic("implement me")
}

func (q QcloudCos) FetchFile(remotePath string, localPath string) (err error) {
	panic("implement me")
}
