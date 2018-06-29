// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/6/29.

package store

import (
	"fmt"
	"time"
	"bytes"
)

func PutPerf(clientId string, timestamp uint64, value []byte) error {
	key := fmt.Sprintf("%s/%s", clientId, time.Unix(int64(timestamp), 0).Format(time.RFC3339))
	return put("performance", key, value)
}

func ListPerf(clientId string) (list [][]byte, err error) {
	tx, err := boltDB.Begin(true)
	if err != nil {
		return nil, err
	}

	bkt := tx.Bucket([]byte("performance"))

	c := bkt.Cursor()
	now := time.Now()
	min := []byte(fmt.Sprintf("%s/%s", clientId, now.Add(-time.Hour*24).Format(time.RFC3339)))
	max := []byte(fmt.Sprintf("%s/%s", clientId, now.Format(time.RFC3339)))

	list = [][]byte{}
	for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Next() {
		list = append(list, v)
	}

	return list, tx.Commit()
}

