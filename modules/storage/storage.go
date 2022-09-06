// Copyright 2019. All rights reserved.
// This file is part of duguying project
// I am coding in Tencent
// Created by rainesli on 2020/4/11.

package storage

import (
	"duguying/studio/g"
	"fmt"
)

type FileInfo struct {
	Path string `json:"path"`
	Size int64  `json:"size"`
	Mime string `json:"mime"`
}

type Storage interface {
	List(remotePrefix string) (list []*FileInfo, err error)
	GetFileInfo(remotePath string) (info *FileInfo, err error)
	PutFile(localPath string, remotePath string) (err error)
	RenameFile(remotePath string, newRemotePath string) (err error)
	RemoveFile(remotePath string) (err error)
	FetchFile(remotePath string, localPath string) (err error)
}

var (
	AliyunCosType = "aliyun"
	QcloudCosType = "qcloud"
)

// NewCos 创建 cos 实例
func NewCos(cosType string) (Storage, error) {
	switch cosType {
	case AliyunCosType:
		{
			ak := g.Config.Get("aliyun-cos", "ak", "")
			sk := g.Config.Get("aliyun-cos", "sk", "")
			bucket := g.Config.Get("aliyun-cos", "bucket", "")
			return NewAliyunOss(ak, sk, bucket)
		}
	case QcloudCosType:
		{
			sid := g.Config.Get("qcloud-cos", "sid", "")
			skey := g.Config.Get("qcloud-cos", "skey", "")
			bucket := g.Config.Get("qcloud-cos", "bucket", "")
			region := g.Config.Get("qcloud-cos", "region", "")
			return NewQcloudOss(*g.LogEntry, sid, skey, bucket, region)
		}
	default:
		{
			return nil, fmt.Errorf("不支持的云存储类型")
		}
	}
}
