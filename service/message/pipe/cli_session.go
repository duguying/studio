// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/6/13.

package pipe

import "github.com/gogather/d2"

var (
	cliSession = d2.NewD2()
)

func SetCliPid(clientId string, reqId string, pid uint32) {
	cliSession.Add(clientId, reqId, pid)
}

func GetCliPid(clientId string, reqId string) (pid uint32, exist bool) {
	val, exist := cliSession.Get(clientId, reqId)
	if exist {
		pid = val.(uint32)
	}
	return pid, exist
}
