// Copyright 2019. All rights reserved.
// This file is part of duguying project
// I am coding in Tencent
// Created by rainesli on 2020/4/11.

package storage

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tencentyun/cos-go-sdk-v5"
)

type QcloudCos struct {
	sid    string
	skey   string
	bucket string
	client *cos.Client
	l      *logrus.Entry
	ctx    context.Context
}

func NewQcloudOss(l *logrus.Entry, sid string, skey string, bucket, region, protocol string) (storage *QcloudCos, err error) {
	storage = &QcloudCos{
		sid:    sid,
		skey:   skey,
		bucket: bucket,
		l:      l,
		ctx:    l.Context,
	}

	u, err := url.Parse(fmt.Sprintf("%s://%s.cos.%s.myqcloud.com", protocol, bucket, region))
	if err != nil {
		return nil, err
	}
	b := &cos.BaseURL{BucketURL: u}
	storage.client = cos.NewClient(b, &http.Client{
		//设置超时时间
		Timeout: 100 * time.Second,
		Transport: &cos.AuthorizationTransport{
			//如实填写账号和密钥，也可以设置为环境变量
			SecretID:  storage.sid,
			SecretKey: storage.skey,
		},
	})

	return storage, nil
}

// List 列举文件
func (q QcloudCos) List(remotePrefix string) (list []*FileInfo, err error) {
	opt := &cos.BucketGetOptions{
		Prefix:  remotePrefix,
		MaxKeys: 3,
	}

	v, _, err := q.client.Bucket.Get(context.Background(), opt)
	if err != nil {
		return nil, err
	}

	list = []*FileInfo{}
	for _, c := range v.Contents {
		list = append(list, &FileInfo{
			Path: c.Key,
			Size: c.Size,
		})
	}

	return list, nil
}

// GetFileInfo 获取文件信息
func (q QcloudCos) GetFileInfo(remotePath string) (info *FileInfo, err error) {
	list, err := q.List(remotePath)
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		info = list[0]
	}
	return info, nil
}

func (q QcloudCos) IsExist(remotePath string) (exist bool, err error) {
	return q.client.Object.IsExist(q.ctx, remotePath)
}

func (q QcloudCos) PutFile(localPath string, remotePath string) (err error) {
	_, err = q.client.Object.PutFromFile(q.ctx, remotePath, localPath, nil)
	if err != nil {
		return err
	}
	return nil
}

func (q QcloudCos) copyFile(remotePath string, newRemotePath string) (err error) {
	_, _, err = q.client.Object.Copy(q.ctx, newRemotePath, remotePath, nil)
	if err != nil {
		return err
	}
	return nil
}

func (q QcloudCos) RenameFile(remotePath string, newRemotePath string) (err error) {
	err = q.copyFile(remotePath, newRemotePath)
	if err != nil {
		return err
	}
	return q.RemoveFile(remotePath)
}

func (q QcloudCos) RemoveFile(remotePath string) (err error) {
	_, err = q.client.Object.Delete(q.ctx, remotePath, nil)
	if err != nil {
		return err
	}
	return nil
}

func (q QcloudCos) FetchFile(remotePath string, localPath string) (err error) {
	_, err = q.client.Object.GetToFile(q.ctx, remotePath, localPath, nil)
	if err != nil {
		return err
	}
	return nil
}
