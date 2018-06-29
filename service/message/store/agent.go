// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/6/29.

package store

import (
	"encoding/json"
	"log"
	"time"
)

type AgentStatusInfo struct {
	Online      bool      `json:"online"`
	ClientID    string    `json:"client_id"`
	Hostname    string    `json:"hostname"`
	Ip          string    `json:"ips"`
	IpIns       []string  `json:"ipins"`
	OnlineTime  time.Time `json:"online_time"`
	OfflineTime time.Time `json:"offline_time"`
}

func (ai *AgentStatusInfo) String() string {
	c, _ := json.Marshal(ai)
	return string(c)
}

func PutAgent(clientId string, info *AgentStatusInfo) error {
	infoOld, err := GetAgent(clientId)
	if err != nil {
		// nop
	} else {
		if info.OfflineTime.IsZero() {
			info.OfflineTime = infoOld.OfflineTime
		}
		if info.OnlineTime.IsZero() {
			info.OnlineTime = infoOld.OnlineTime
		}
		if info.Hostname == "" {
			info.Hostname = infoOld.Hostname
		}
		if info.Ip == "" {
			info.Ip = infoOld.Ip
		}
		if len(info.IpIns) <= 0 || info.IpIns[0] == "" {
			info.IpIns = infoOld.IpIns
		}
	}

	value := info.String()

	return put("agent", clientId, []byte(value))
}

func GetAgent(clientId string) (info *AgentStatusInfo, err error) {
	tx, err := boltDB.Begin(true)
	if err != nil {
		return nil, err
	}

	bkt := tx.Bucket([]byte("agent"))

	value := bkt.Get([]byte(clientId))
	info = &AgentStatusInfo{}
	err = json.Unmarshal(value, info)
	if err != nil {
		return nil, err
	}
	return info, tx.Commit()
}

func ListAllAgent() (list []*AgentStatusInfo, err error) {
	tx, err := boltDB.Begin(true)
	if err != nil {
		return nil, err
	}

	bkt := tx.Bucket([]byte("agent"))

	c := bkt.Cursor()
	list = []*AgentStatusInfo{}

	for k, v := c.First(); k != nil; k, v = c.Next() {
		info := &AgentStatusInfo{}
		err := json.Unmarshal(v, info)
		if err != nil {
			log.Println("marshal agent info failed, err:", err.Error())
		} else {
			list = append(list, info)
		}
	}

	return list, tx.Commit()
}

func clearAllAgentPerf() (err error) {
	list, err := ListAllAgent()
	if err != nil {
		return err
	}
	for _, agent := range list {
		return clearRange(agent.ClientID)
	}
	return nil
}
