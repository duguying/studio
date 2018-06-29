// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/6/29.

package store

import (
	"bytes"
	"duguying/studio/service/message/model"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"time"
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
		info := &AgentStatusInfo{
			ClientID: clientId,
			IpIns:    ips,
			Hostname: perf.Hostname,
		}
		PutAgent(clientId, info)
	}

	key := fmt.Sprintf("%s/%s", clientId, time.Unix(int64(timestamp), 0).Format(time.RFC3339))
	return put("performance", key, value)
}

func ListPerf(clientId string) (list []*model.PerformanceMonitor, err error) {
	tx, err := boltDB.Begin(true)
	if err != nil {
		return nil, err
	}

	bkt := tx.Bucket([]byte("performance"))

	c := bkt.Cursor()
	now := time.Now()
	min := []byte(fmt.Sprintf("%s/%s", clientId, now.Add(-time.Hour*24).Format(time.RFC3339)))
	max := []byte(fmt.Sprintf("%s/%s", clientId, now.Format(time.RFC3339)))

	list = []*model.PerformanceMonitor{}
	for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Next() {
		perf := &model.PerformanceMonitor{}
		err := proto.Unmarshal(v, perf)
		if err != nil {
			continue
		}
		list = append(list, perf)
	}

	return list, tx.Commit()
}

func clearRange(clientId string) (err error) {
	tx, err := boltDB.Begin(true)
	if err != nil {
		return err
	}

	bkt := tx.Bucket([]byte("performance"))

	c := bkt.Cursor()
	now := time.Now()
	min := []byte(fmt.Sprintf("%s/%s", clientId, time.Now().Add(-time.Hour*24*365).Format(time.RFC3339)))
	max := []byte(fmt.Sprintf("%s/%s", clientId, now.Add(-time.Hour*24).Format(time.RFC3339)))
	for k, _ := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, _ = c.Next() {
		err := bkt.Delete(k)
		if err != nil {
			log.Println("delete err:", err.Error())
		} else {
			log.Println("delete key:", string(k))
		}
	}

	return tx.Commit()
}
