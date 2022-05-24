// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/6/13.

package pipe

import (
	"fmt"
	"github.com/gogather/d2"
)

var (
	d2map = d2.NewD2()
)

type ChanPair struct {
	ChanIn  chan []byte
	ChanOut chan []byte
}

func NewCliChanPair() (pair *ChanPair) {
	return &ChanPair{
		ChanIn:  make(chan []byte, 10000),
		ChanOut: make(chan []byte, 10000),
	}
}

func SetCliChanPair(session string, pid uint32, pair *ChanPair) {
	d2map.Add(session, fmt.Sprintf("%d", pid), pair)
}

func GetCliChanPair(session string, pid uint32) (pair *ChanPair, exist bool) {
	val, exist := d2map.Get(session, fmt.Sprintf("%d", pid))
	if exist {
		pair = val.(*ChanPair)
	}
	return pair, exist
}
