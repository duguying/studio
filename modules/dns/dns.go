// Copyright 2019. All rights reserved.
// This file is part of duguying project
// I am coding in Tencent
// Created by rainesli on 2020/4/19.

package dns

const (
	ARecord     = "A"
	AAAARecord  = "AAAA"
	CnameRecord = "CNAME"
	TxtRecord   = "TXT"
)

type Dns interface {
	AddDomainRecord(domainName string, recordType string, record string, value string) (recordId string, err error)
	DeleteDomainRecord(recordId string) (err error)
}
