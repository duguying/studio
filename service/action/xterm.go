// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/6/13.

package action

import (
	"duguying/studio/service/message/model"
	"duguying/studio/service/message/pipe"
	"duguying/studio/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

func ConnectXTerm(c *gin.Context) {
	clientId := c.Query("client_id")

	if clientId == "" {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": "client_id is required",
		})
		return
	}

	// create cli
	reqId, err := utils.GenUUID()
	if err != nil {
		log.Println("generate uuid failed:", err)
		c.JSON(http.StatusForbidden, map[string]interface{}{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}
	openCliCmd := model.CliCmd{
		Cmd:       model.CliCmd_OPEN,
		Session:   clientId,
		RequestId: reqId,
		Pid:       0,
	}
	pcmdData, err := proto.Marshal(&openCliCmd)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}
	success := pipe.SendMsg(clientId, model.Msg{
		Type:     websocket.BinaryMessage,
		Cmd:      model.CMD_CLI_CMD,
		ClientId: clientId,
		Data:     pcmdData,
	})
	if !success {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": "created cli cmd send failed",
		})
		return
	}

	// wait creation cli response and get pid
	pid := uint32(0)
	for i := 0; i < 10000; i++ {
		time.Sleep(time.Millisecond)
		var exist = false
		pid, exist = pipe.GetCliPid(clientId, reqId)
		if exist {
			break
		}
	}

	if pid <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": "invalid pid, maybe create cli failed",
		})
		return
	}

	// upgrade to websocket
	var upgrader = websocket.Upgrader{}
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, c.Writer.Header())
	if err != nil {
		log.Println("upgrade:", err)
		c.JSON(http.StatusForbidden, map[string]interface{}{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	defer conn.Close()

	pair := pipe.NewCliChanPair()
	pipe.SetCliChanPair(clientId, pid, pair)

	// send xterm data into cli
	go func() {
		for {
			select {
			case data := <-pair.ChanOut:
				{
					pipeStruct := model.CliPipe{
						Session: clientId,
						Pid:     pid,
						Data:    data,
					}
					pipeData, err := proto.Marshal(&pipeStruct)
					if err != nil {
						log.Println("proto marshal failed, err:", err.Error())
						continue
					}
					msg := model.Msg{
						Type:     websocket.BinaryMessage,
						Cmd:      model.CMD_CLI_PIPE,
						ClientId: clientId,
						Data:     pipeData,
					}
					pipe.SendMsg(clientId, msg)
				}
			}
		}
	}()

	// read from client, put into in channel
	go func(con *websocket.Conn) {
		for {
			_, data, err := con.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}

			pair.ChanOut <- data
		}
	}(conn)

	// write into client, get from out channel
	for {
		select {
		case data := <-pair.ChanIn:
			{
				err = conn.WriteMessage(websocket.TextMessage, data)
				if err != nil {
					log.Println("即时消息发送到客户端:", err)
					break
				}
			}
		}
	}

}
