// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/6/7.

package deal

import (
	"duguying/blog/service/message/model"
	"duguying/blog/service/message/pipe"
	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

func InitHb() {
	go func() {
		for {
			sendHb()
			time.Sleep(time.Second * 30)
		}
	}()
}

func sendHb() {
	cm := pipe.GetConMap()
	if cm == nil {
		return
	}

	hbp := &model.HeartBeat{
		Timestamp: uint64(time.Now().Unix()),
	}
	packet, err := proto.Marshal(hbp)
	if err != nil {
		log.Println("marshal proto failed, err:", err.Error())
		return
	}

	for clientId, _ := range cm.M {
		pipe.SendMsg(clientId, model.Msg{
			Type: websocket.BinaryMessage,
			Cmd:  model.CMD_HB,
			Data: packet,
		})
	}
}
