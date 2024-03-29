// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/6/7.

package agent

import (
	"duguying/studio/g"
	"duguying/studio/modules/db"
	"duguying/studio/service/message/model"
	"duguying/studio/service/message/pipe"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gogather/com"
	"github.com/gorilla/websocket"
)

func Ws(c *gin.Context) {
	clientId := c.Query("client_id")

	if clientId == "" {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": "client_id is required",
		})
		return
	}

	log.Println("ws connect with client_id:", clientId)

	var upgrader = websocket.Upgrader{}
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, c.Writer.Header())
	if err != nil {
		// 已经 upgrade 为 websocket，不能再按 http 写入
		log.Println("upgrade:", err)
		return
	}

	// client ip
	ip := c.ClientIP()

	// store connects
	connId := com.CreateGUID()
	pipe.AddConnect(connId, conn)

	defer conn.Close()
	out := make(chan model.Msg, 100)

	// register in and out channel
	pipe.AddUserPipe(clientId, out, connId)

	// store agent info
	_, err = db.CreateOrUpdateAgent(g.Db, clientId, ip)
	if err != nil {
		log.Println("put agent failed, err:", err.Error())
	}

	// read from client, put into in channel
	go func(con *websocket.Conn) {
		for {
			var err error

			mt, msgData, err := con.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}

			msg := model.Msg{
				Type:     mt,
				Cmd:      int(msgData[0]),
				ClientId: clientId,
				Data:     msgData[1:],
			}

			if g.Config.Get("ws", "log", "enable") == "enable" {
				log.Printf("recv: %s\n", msg.Info())
			}

			pipe.In <- msg
		}
	}(conn)

	// write into client, get from out channel
	for {
		var err error
		var msg model.Msg

		msg = <-out
		//log.Println("send message:", msg.String())

		err = conn.WriteMessage(msg.Type, append([]byte{byte(msg.Cmd)}, msg.Data...))
		if err != nil {
			log.Println("即时消息发送到客户端:", err)
			break
		}
	}

	// exit websocket finally, and remove client pipeline
	pipe.RemoveUserPipe(clientId)
	pipe.RemoveConnect(connId)

	// update agent info
	err = db.UpdateAgentOffline(g.Db, clientId)
	if err != nil {
		log.Println("update agent offline failed, err:", err.Error())
	}
}
