// Copyright 2020. All rights reserved.
// This file is part of duguying project
// I am coding in Tencent
// Created by rainesli on 2020/3/16.

package viewcnt

import (
	"github.com/gogather/safemap"
)

var (
	viewCntMap = safemap.New()
)

func ViewHit(ident string) {
	val, ok := viewCntMap.Get(ident)
	if !ok {
		viewCntMap.Put(ident, int(1))
	} else {
		cnt := val.(int)
		viewCntMap.Put(ident, cnt+1)
	}
}

func GetViewCnt(ident string) (cnt int) {
	val, ok := viewCntMap.Get(ident)
	if ok {
		return val.(int)
	} else {
		return 0
	}
}

func ResetViewCnt(ident string) {
	viewCntMap.Remove(ident)
}

func GetMap() *safemap.SafeMap {
	return viewCntMap
}
