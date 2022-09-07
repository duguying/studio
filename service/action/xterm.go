// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/6/13.

package action

import (
	"duguying/studio/g"
	"duguying/studio/service/message/model"
	"duguying/studio/service/message/pipe"
	"duguying/studio/utils"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gogather/json"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

type TermLayout struct {
	Width  uint32 `json:"cols"`
	Height uint32 `json:"rows"`
}

func (tl *TermLayout) String() string {
	c, _ := json.Marshal(tl)
	return string(c)
}

func ConnectXTerm(c *gin.Context) {
	clientID := c.Query("client_id")

	if clientID == "" {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": "client_id is required",
		})
		return
	}

	// create cli
	reqID := utils.GenUUID()
	openCliCmd := model.CliCmd{
		Cmd:       model.CliCmd_OPEN,
		Session:   clientID,
		RequestId: reqID,
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
	success := pipe.SendMsg(clientID, model.Msg{
		Type:     websocket.BinaryMessage,
		Cmd:      model.CMD_CLI_CMD,
		ClientId: clientID,
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
		pid, exist = pipe.GetCliPid(clientID, reqID)
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

	wsExit := false
	defer conn.Close()
	defer func() { wsExit = true }()

	pair := pipe.NewCliChanPair()
	pipe.SetCliChanPair(clientID, pid, pair)
	pipe.SetPidCon(clientID, pid, conn) // store connection

	// send xterm data into cli
	go func() {
		for {
			select {
			case data := <-pair.ChanOut:
				{
					if wsExit {
						return
					}
					pipeStruct := model.CliPipe{
						Session: clientID,
						Pid:     pid,
						Data:    data,
					}
					g.LogEntry.WithField("slice", "agentrcv").Printf("agent received in: %s\n", string(data))
					pipeData, err := proto.Marshal(&pipeStruct)
					if err != nil {
						log.Println("proto marshal failed, err:", err.Error())
						continue
					}
					msg := model.Msg{
						Type:     websocket.BinaryMessage,
						Cmd:      model.CMD_CLI_PIPE,
						ClientId: clientID,
						Data:     pipeData,
					}
					pipe.SendMsg(clientID, msg)
				}
			}
		}
	}()

	// read from client, put into in channel
	go func(con *websocket.Conn) {
		for {
			_, data, err := con.ReadMessage()
			if err != nil {
				// ws has closed
				log.Println("read:", err)

				// try to send close cmd to agent cli
				_, exist := pipe.GetPidCon(clientID, pid)
				if exist {
					cliCmdStruct := model.CliCmd{
						Cmd:       model.CliCmd_CLOSE,
						Session:   clientID,
						RequestId: reqID,
						Pid:       pid,
					}
					cliCmdData, err := proto.Marshal(&cliCmdStruct)
					if err != nil {
						log.Println("marshal cmd data failed, err:", err.Error())
					} else {
						cmdCloseMsg := model.Msg{
							Type:     websocket.BinaryMessage,
							Cmd:      model.CMD_CLI_CMD,
							ClientId: clientID,
							Data:     cliCmdData,
						}
						pipe.SendMsg(clientID, cmdCloseMsg)
					}
				}
				break
			}

			if data[0] == model.TERM_PONG {
				//log.Println("pong")
			} else if data[0] == model.TERM_SIZE {
				layout := TermLayout{}
				err = json.Unmarshal(data[1:], &layout)
				if err != nil {
					log.Printf("parse layout failed, err: %s, raw content is: %s\n", err.Error(), string(data[1:]))
					continue
				}
				log.Println("resize...", layout)
				cliCmdStruct := model.CliCmd{
					Cmd:       model.CliCmd_RESIZE,
					Session:   clientID,
					RequestId: reqID,
					Pid:       pid,
					Width:     layout.Width,
					Height:    layout.Height,
				}
				cliCmdData, err := proto.Marshal(&cliCmdStruct)
				if err != nil {
					log.Println("marshal cmd data failed, err:", err.Error())
				} else {
					cmdCloseMsg := model.Msg{
						Type:     websocket.BinaryMessage,
						Cmd:      model.CMD_CLI_CMD,
						ClientId: clientID,
						Data:     cliCmdData,
					}
					pipe.SendMsg(clientID, cmdCloseMsg)
				}
			} else if data[0] == model.TERM_PIPE {
				//log.Printf("what's header: %d\n", data[0])
				g.LogEntry.WithField("slice", "browsersnt").Printf("browser sent: %s\n", string(data[1:]))
				pair.ChanOut <- data[1:]
			}

		}
	}(conn)

	// send hb to xterm
	go func() {
		xtermHbPeriod := g.Config.GetInt64("xterm", "hb", 10)
		for {
			if wsExit {
				return
			}
			pair.ChanIn <- []byte{model.TERM_PING}
			time.Sleep(time.Second * time.Duration(xtermHbPeriod))
		}
	}()

	var wsLock sync.Mutex
	// write into client, get from out channel
	for {
		select {
		case data := <-pair.ChanIn:
			{
				wsLock.Lock()
				g.LogEntry.WithField("slice", "browserrcv").Printf("browser received: len --> %d\n", len(data))
				err = conn.WriteMessage(websocket.BinaryMessage, data)
				if err != nil {
					log.Println("即时消息发送到客户端:", err)
					return
				}
				wsLock.Unlock()
			}
		}
	}
}
