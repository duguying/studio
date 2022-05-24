// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/6/13.

package pipe

import (
	"github.com/gogather/d2"
	"github.com/gorilla/websocket"
	"fmt"
)

var (
	cliSession = d2.NewD2()
	conSession = d2.NewD2()
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

func DelCliPid(clientId string, reqId string) {
	cliSession.RemoveKey(clientId, reqId)
}

// --------

func SetPidCon(clientId string, pid uint32, conn *websocket.Conn) {
	conSession.Add(clientId, fmt.Sprintf("%d",pid), conn)
}

func GetPidCon(clientId string, pid uint32) (conn *websocket.Conn, exist bool) {
	val, exist := conSession.Get(clientId, fmt.Sprintf("%d",pid))
	if exist {
		conn = val.(*websocket.Conn)
	}
	return conn, exist
}

func DelPidCon(clientId string, pid uint32) {
	conSession.RemoveKey(clientId, fmt.Sprintf("%d", pid))
}
