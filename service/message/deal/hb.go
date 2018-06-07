// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/6/7.

package deal

import (
	"duguying/blog/service/message/pipe"
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

	for clientId, _ := range cm.M {
		pipe.SendMsg(clientId, "")
	}
}
