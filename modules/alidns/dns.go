// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/5/18.

package alidns

import (
	"duguying/blog/utils"
	"encoding/json"
	"fmt"
	"github.com/Unknwon/com"
	"github.com/parnurzeal/gorequest"
	"log"
	"net/http"
	"sort"
	"time"
)

type AliArgsBase struct {
	Format           string `json:"Format"`
	Version          string `json:"version"`
	SignatureMethod  string `json:"SignatureMethod"`
	SignatureNonce   string `json:"SignatureNonce"`
	SignatureVersion string `json:"SignatureVersion"`
	AccessKeyId      string `json:"AccessKeyId"`
	Sk               string `json:"-"`
	RootAddr         string `json:"-"`
	Timestamp        string `json:"Timestamp"`
}

func NewAliArgsBase(host string, ak, sk string) AliArgsBase {
	guid, _ := utils.GenUUID()
	return AliArgsBase{
		Format:           "json",
		Version:          "2015-01-09",
		SignatureMethod:  "HMAC-SHA1",
		SignatureNonce:   guid,
		SignatureVersion: "1.0",
		AccessKeyId:      ak,
		Sk:               sk,
		RootAddr:         host,
		Timestamp:        time.Now().UTC().Format("2006-01-02T15:04:05Z"),
	}
}

type AliAddRecord struct {
	AliArgsBase

	Action     string `json:"Action"`
	DomainName string `json:"DomainName"`
	RR         string `json:"RR"`
	Type       string `json:"Type"`
	Value      string `json:"Value"`
	TTL        string `json:"TTL"`
	Line       string `json:"Line"`
}

func (aar *AliAddRecord) String() string {
	c, _ := json.Marshal(aar)
	return string(c)
}

func (arr *AliAddRecord) ToURL() string {
	args := map[string]string{}
	json.Unmarshal([]byte(arr.String()), &args)

	keys := []string{}
	for key, _ := range args {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	argsLine := ""
	for _, key := range keys {
		argsLine = argsLine + fmt.Sprintf("&%s=%s", key, args[key])
	}

	stringToSign := fmt.Sprintf("GET&/%s", argsLine)
	stringToSign = com.UrlEncode(stringToSign)
	signature := utils.HmacSha1(stringToSign, fmt.Sprintf("%s&", arr.Sk))

	addr := fmt.Sprintf("%s/?%s", arr.RootAddr, fmt.Sprintf("%s%s", fmt.Sprintf("Signature=%s", signature), argsLine))
	return addr
}

func NewAliAddRecord(host, ak, sk string, rootDomain string, RR string, recordType string, ttl int, line string, value string) AliAddRecord {
	return AliAddRecord{
		Action:     "AddDomainRecord",
		DomainName: rootDomain,
		RR:         RR,
		Type:       recordType,
		TTL:        fmt.Sprintf("%d", ttl),
		Line:       value,

		AliArgsBase: NewAliArgsBase(host, ak, sk),
	}
}

func AddDomainRecord(host, ak, sk string, rootDomain string, RR string, recordType string, ttl int, line string, value string) (err error) {
	entity := NewAliAddRecord(host, ak, sk, rootDomain, RR, recordType, 60, line, value)
	link := entity.ToURL()
	log.Println("add domain request:", link)
	resp, content, errs := gorequest.New().Get(link).End()
	if len(errs) > 0 && errs[0] != nil {
		return errs[0]
	} else if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid http status: %s", resp.Status)
	} else {
		log.Println(content)
	}

	return
}
