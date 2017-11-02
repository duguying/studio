// Copyright 2017. All rights reserved.
// This file is part of cleaner project
// Created by duguying on 2017/10/1.

package cleaner

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

type FCleaner struct {
	root     string
	expire   time.Duration
	interval time.Duration
	batchNum int64
	filter   func(info os.FileInfo) (willClean bool)
	count    int64
}

// New new a Cleaner
// root - root directory
// expire - expire time
// interval - scan interval
// batchNum - every scan remove batch number limit
func New(root string, expire time.Duration, interval time.Duration, batchNum int64) *FCleaner {
	return &FCleaner{
		root:     root,
		expire:   expire,
		interval: interval,
		batchNum: batchNum,
		filter: func(info os.FileInfo) (willClean bool) {
			return true
		},
	}
}

func (fc *FCleaner) SetFilter(filter func(info os.FileInfo) (willClean bool)) {
	fc.filter = filter
}

// StartCleanTask start the cleaning task
func (fc *FCleaner) StartCleanTask() {
	go func(fc *FCleaner) {
		for {
			err := fc.clean(fc.root)
			if err != nil {
				log.Printf("[FCleaner] clean err: %v\n", err)
			}
			log.Printf("[FCleaner] total file number: %d\n", fc.count)
			time.Sleep(fc.interval)
		}
	}(fc)
}

func (fc *FCleaner) clean(root string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if path == root {
			return nil
		}

		if fc.batchNum > 0 && fc.count > fc.batchNum {
			return nil
		}

		if info.IsDir() {
			// do sth
			err := fc.clean(path)
			if err != nil {
				return err
			} else {
				empty, err := fc.isEmpty(path)
				if err == nil && empty {
					// remove empty directory
					fc.removePath(path)
				}
				return nil
			}
		} else {
			// do clean
			if fc.filter(info) && fc.checkExpire(info) {
				fc.removePath(path)
			}
			return nil
		}
	})
}

func (fc *FCleaner) removePath(path string) error {
	fc.count++
	fmt.Println(path)
	return os.Remove(path)
}

func (fc *FCleaner) checkExpire(info os.FileInfo) (isExpire bool) {
	if info.ModTime().Add(fc.expire).Before(time.Now()) {
		return true
	} else {
		return false
	}
}

func (fc *FCleaner) isEmpty(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}
