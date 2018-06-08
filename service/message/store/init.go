// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/6/8.

package store

import (
	"bytes"
	"duguying/blog/g"
	"github.com/boltdb/bolt"
	"log"
	"time"
)

var (
	boltDB *bolt.DB
)

func InitBoltDB() {
	dbPath := g.Config.Get("boltdb", "path", "performance.db")
	// open db
	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		boltDB = db
	}
}

func Put(timestamp uint64, value []byte) error {
	tx, err := boltDB.Begin(true)
	if err != nil {
		return err
	}

	b, err := tx.CreateBucketIfNotExists([]byte("performance"))
	if err != nil {
		return err
	}

	err = b.Put([]byte(time.Unix(int64(timestamp), 0).Format(time.RFC3339)), value)
	if err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}

func List() (list [][]byte, err error) {
	tx, err := boltDB.Begin(true)
	if err != nil {
		return nil, err
	}

	bkt := tx.Bucket([]byte("performance"))

	c := bkt.Cursor()
	now := time.Now()
	min := []byte(now.Add(-time.Hour * 24).Format(time.RFC3339))
	max := []byte(now.Format(time.RFC3339))

	list = [][]byte{}
	for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Next() {
		list = append(list, v)
	}

	return list, tx.Commit()
}
