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

	qcos "github.com/tencentyun/cos-go-sdk-v5"
)

// QCloudCos 腾讯云 COS 存储
type QCloudCos struct {
	ak         string
	sk         string
	bucket     string
	appID      string
	region     string
	bucketHost string
	client     *qcos.Client
}

// NewQCloudCos 新建 QCloudCos
func NewQCloudCos(ak string, sk string, bucket string, appID string, region string) (storage *QCloudCos, err error) {
	storage = &QCloudCos{
		ak:     ak,
		sk:     sk,
		bucket: bucket,
		appID:  appID,
		region: region,
	}

	// 将 examplebucket-1250000000 和 COS_REGION 修改为用户真实的信息
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。https://console.cloud.tencent.com/cos5/bucket
	// COS_REGION 可以在控制台查看，https://console.cloud.tencent.com/cos5/bucket, 关于地域的详情见 https://cloud.tencent.com/document/product/436/6224
	u, err := url.Parse(fmt.Sprintf("https://%s-%s.cos.%s.myqcloud.com", storage.bucket, storage.appID, storage.region))
	if err != nil {
		return nil, err
	}
	storage.bucketHost = u.Host

	// 用于Get Service 查询，默认全地域 service.cos.myqcloud.com
	su, err := url.Parse(fmt.Sprintf("https://cos.%s.myqcloud.com", storage.region))
	if err != nil {
		return nil, err
	}

	b := &qcos.BaseURL{
		BucketURL:  u,
		ServiceURL: su,
	}

	// 创建 client
	client := qcos.NewClient(b, &http.Client{
		Transport: &qcos.AuthorizationTransport{
			SecretID:  "",
			SecretKey: "",
		},
	})
	storage.client = client

	return storage, nil
}

// List 列举文件
func (q *QCloudCos) List(remotePrefix string) (list []*FileInfo, err error) {
	opt := &qcos.BucketGetOptions{
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
func (q *QCloudCos) GetFileInfo(remotePath string) (info *FileInfo, err error) {
	list, err := q.List(remotePath)
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		info = list[0]
	}
	return info, nil
}

// AddFile 向存储桶上传文件
func (q *QCloudCos) AddFile(localPath string, remotePath string) (err error) {
	// 通过本地文件上传对象
	_, err = q.client.Object.PutFromFile(context.Background(), remotePath, localPath, nil)
	if err != nil {
		return err
	}
	return nil
}

// RenameFile 重命名文件
func (q *QCloudCos) RenameFile(sourceRemotePath string, destRemotePath string) (err error) {
	sourceURL := fmt.Sprintf("%s/%s", q.bucketHost, sourceRemotePath)
	_, _, err = q.client.Object.Copy(context.Background(), destRemotePath, sourceURL, nil)
	if err != nil {
		return err
	}
	_, err = q.client.Object.Delete(context.Background(), sourceRemotePath, nil)
	if err != nil {
		return err
	}
	return nil
}

// RemoveFile 删除文件
func (q *QCloudCos) RemoveFile(remotePath string) (err error) {
	_, err = q.client.Object.Delete(context.Background(), remotePath)
	if err != nil {
		return err
	}
	return nil
}

// FetchFile 获取文件（下载）
func (q *QCloudCos) FetchFile(remotePath string, localPath string) (err error) {
	opt := &qcos.MultiDownloadOptions{
		ThreadPoolSize: 5,
	}
	_, err = q.client.Object.Download(
		context.Background(), remotePath, localPath, opt,
	)
	if err != nil {
		return err
	}
	return nil
}
