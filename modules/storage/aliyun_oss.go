// Copyright 2019. All rights reserved.
// This file is part of duguying project
// I am coding in Tencent
// Created by rainesli on 2020/4/11.

package storage

type AliyunOss struct {
	ak     string
	sk     string
	bucket string
}

func NewAliyunOss(ak string, sk string, bucket string) (storage *AliyunOss, err error) {
	storage = &AliyunOss{
		ak:     ak,
		sk:     sk,
		bucket: bucket,
	}
	return storage, nil
}

func (AliyunOss) List(remotePrefix string) (list []*FileInfo, err error) {
	panic("implement me")
}

func (AliyunOss) GetFileInfo(remotePath string) (info *FileInfo, err error) {
	panic("implement me")
}

func (AliyunOss) IsExist(remotePath string) (exist bool, err error) {
	panic("implement me")
}

func (AliyunOss) PutFile(localPath string, remotePath string) (err error) {
	panic("implement me")
}

func (AliyunOss) RenameFile(remotePath string, newRemotePath string) (err error) {
	panic("implement me")
}

func (AliyunOss) RemoveFile(remotePath string) (err error) {
	panic("implement me")
}

func (AliyunOss) FetchFile(remotePath string, localPath string) (err error) {
	panic("implement me")
}
