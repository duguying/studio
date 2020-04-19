// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/5/18.

package dns

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

type AliDns struct {
	ak     string
	sk     string
	client *alidns.Client
}

func NewAliDns(ak string, sk string) (cli *AliDns, err error) {
	client, err := alidns.NewClient()
	if err != nil {
		return nil, err
	}
	return &AliDns{
		ak:     ak,
		sk:     sk,
		client: client,
	}, nil
}

func (ad *AliDns) AddDomainRecord(domainName string, recordType string, record string, value string) (recordId string, err error) {
	response, err := ad.client.AddDomainRecord(&alidns.AddDomainRecordRequest{
		RR:         record,
		Type:       recordType,
		Value:      value,
		DomainName: domainName,
	})
	if err != nil {
		return "", err
	}
	if !response.IsSuccess() {
		return "", fmt.Errorf("添加失败")
	}
	return response.RecordId, nil
}

func (ad *AliDns) DeleteDomainRecord(recordId string) (err error) {
	response, err := ad.client.DeleteDomainRecord(&alidns.DeleteDomainRecordRequest{
		RecordId: recordId,
	})
	if err != nil {
		return err
	}
	if !response.IsSuccess() {
		return fmt.Errorf("删除失败")
	}
	return nil
}
