// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/6/29.

package store

import (
	"duguying/studio/service/db"
	"duguying/studio/service/message/model"
	"log"

	"github.com/golang/protobuf/proto"
)

func PutPerf(clientId string, timestamp uint64, value []byte) error {
	perf := &model.PerformanceMonitor{}
	err := proto.Unmarshal(value, perf)
	if err != nil {
		log.Println("proto unmarshal failed, err:", err.Error())
	} else {
		ips := []string{}
		for _, network := range perf.Nets {
			ips = append(ips, network.Ip)
		}
		err = db.PutPerf(clientId, perf.Os, perf.Arch, perf.Hostname, ips)
		if err != nil {
			log.Println("put agent failed, err:", err.Error())
		}
	}

	return nil
}
