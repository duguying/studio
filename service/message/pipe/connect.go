// Copyright 2017. All rights reserved.
// This file is part of im project
// Created by duguying on 2017/9/29.

package pipe

import (
	"github.com/gogather/safemap"
	"github.com/gorilla/websocket"
)

var conns *safemap.SafeMap

func GetConnMap() *safemap.SafeMap {
	return conns
}

func AddConnect(connId string, conn *websocket.Conn) {
	conns.Put(connId, conn)
}

func RemoveConnect(connId string) {
	conns.Remove(connId)
}

func GetConnect(connId string) (*websocket.Conn, bool) {
	connect, exist := conns.Get(connId)
	return connect.(*websocket.Conn), exist
}
