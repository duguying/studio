// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/6/7.

package model

import (
	"fmt"
	"log"

	"github.com/gogather/json"
	"github.com/gorilla/websocket"
)

const (
	CMD_HB       = 0
	CMD_PERF     = 1
	CMD_KEY      = 2
	CMD_CLI_PIPE = 3
	CMD_CLI_CMD  = 4

	TERM_PIPE = 0x00
	TERM_SIZE = 0x01
	TERM_PING = 0x02
	TERM_PONG = 0x03
)

type Msg struct {
	Type     int    `json:"type"`
	Cmd      int    `json:"cmd"`
	ClientId string `json:"client_id"`
	Data     []byte `json:"data"`
}

func (m *Msg) String() string {
	ds := map[string]interface{}{
		"type":      m.Type,
		"cmd":       m.Cmd,
		"client_id": m.ClientId,
	}
	if m.Type == websocket.TextMessage {
		ds["data"] = fmt.Sprintf("%s", string(m.Data))
	} else {
		ds["data"] = fmt.Sprintf("%v", m.Data)
	}
	c, err := json.Marshal(ds)
	if err != nil {
		log.Println("json marshal failed, err:", err.Error())
	}
	return string(c)
}

func (m *Msg) Info() string {
	ds := map[string]interface{}{
		"type":      m.Type,
		"cmd":       m.Cmd,
		"client_id": m.ClientId,
		"_data_len": len(m.Data),
	}
	c, err := json.Marshal(ds)
	if err != nil {
		log.Println("json marshal failed, err:", err.Error())
	}
	return string(c)
}
