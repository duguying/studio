// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/5/18.

package dns

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestAddDomainRecord(t *testing.T) {
	client, err := NewAliDns("bw4tirpW2iUODxRI", "uqoJmlNeeUnoBfPJda6OaSj1pLpTPD")
	if err != nil {
		log.Println("client failed, err:", err.Error())
		return
	}
	recordId, err := client.AddDomainRecord("duguying.net", ARecord, "rpi", "127.0.0.1")
	if err != nil {
		log.Println("add failed, err:", err.Error())
		return
	}
	fmt.Println("success, recordId:", recordId)

	time.Sleep(time.Minute * 10)

	err = client.DeleteDomainRecord(recordId)
	if err != nil {
		log.Println("delete failed, err:", err.Error())
		return
	}
}
