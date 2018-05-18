// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/5/18.

package alidns

import (
	"testing"
	"fmt"
)

func TestAddDomainRecord(t *testing.T) {
	base := NewAliArgsBase()
	arg := AliAddRecord{}

	arg.AliArgsBase = base

	fmt.Println(arg.String())

	//arg.Format = "json"
	//arg.Version = "2015-01-09"
	//arg.SignatureMethod = "HMAC-SHA1"
	//arg.SignatureNonce = guid
	//arg.SignatureVersion = "1.0"
	//arg.AccessKeyId = ak
	//arg.Timestamp = time.Now().UTC().Format("2006-01-02T15:04:05Z")


}
