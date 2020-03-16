// Copyright 2020. All rights reserved.
// This file is part of duguying project
// I am coding in Tencent
// Created by rainesli on 2020/3/16.

package viewcnt

import (
	"fmt"
	"github.com/gogather/safemap"
)

var (
	viewCntMap = safemap.New()
)

func ViewHit(articleId int) {
	key := fmt.Sprintf("%d", articleId)
	val, ok := viewCntMap.Get(key)
	if !ok {
		viewCntMap.Put(key, int(1))
	} else {
		cnt := val.(int)
		viewCntMap.Put(key, cnt+1)
	}
}

func GetViewCnt(articleId int) (cnt int) {
	key := fmt.Sprintf("%d", articleId)
	val, ok := viewCntMap.Get(key)
	if ok {
		return val.(int)
	} else {
		return 0
	}
}

func ResetViewCnt(articleId int) {
	key := fmt.Sprintf("%d", articleId)
	viewCntMap.Remove(key)
}

func GetMap() *safemap.SafeMap {
	return viewCntMap
}
