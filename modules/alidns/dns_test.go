// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/5/18.

package alidns

import (
	"fmt"
	"testing"
)

func TestAddDomainRecord(t *testing.T) {
	rcd := NewAliAddRecord("", "ak", "sk", "duguying.net", "rpi", "A", 60, "default", "127.0.0.1")
	fmt.Println(rcd.ToURL())
}
