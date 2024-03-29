// Copyright 2017. All rights reserved.
// This file is part of ofs project
// Created by duguying on 2017/11/29.

package cache

import (
	"fmt"
	"log"
	"sort"
	"testing"
	"time"

	"gopkg.in/redis.v5"
)

var (
	redisCli *redis.Client
)

func initRedis() {
	readTimeout := 4
	db := 2
	redisCli = redis.NewClient(&redis.Options{
		Addr:        "127.0.0.1:6379",
		Password:    "",
		DB:          db,
		PoolSize:    10000,
		ReadTimeout: time.Duration(time.Second * time.Duration(readTimeout)),
	})
	err := redisCli.Ping().Err()

	if err != nil {
		log.Println("[system]", err.Error())
	} else {
		log.Println("[system]", "redis connect success")
	}
}

func TestSetTTL(t *testing.T) {
	initRedis()

	//err := SetMapField("hi", "hello", "1244")
	//fmt.Println(err)
	//
	//fmt.Println(GetMap("hi"))
	//DelMapField("hi","hello")

	redisCli.Set("hi", "hello", 0)
	fmt.Println(redisCli.Get("hi"))
}

func TestGet(t *testing.T) {
	args := []string{"casdf", "badfadf", "basd", "a"}
	fmt.Println(args)

	sort.Strings(args)
	fmt.Println(args)
}
