// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/5/18.

package agent

import (
	"duguying/studio/g"
	"duguying/studio/modules/alidns"
	"duguying/studio/modules/db"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AgentInfo struct {
	Ips    []string `json:"ips"`
	CpuNum int      `json:"cpu_num"`
}

func Report(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	authToken := g.Config.Get("dns", "agent-auth", "a466e30d7571e6e720cb4a01ce446752")
	if auth != authToken {
		c.JSON(http.StatusUnauthorized, gin.H{
			"ok":  false,
			"err": "auth failed",
		})
	}

	ai := &AgentInfo{}
	err := c.BindJSON(ai)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	rootAddr := g.Config.Get("dns", "addr", "http://alidns.aliyuncs.com")
	ak := g.Config.Get("dns", "ak", "")
	sk := g.Config.Get("dns", "sk", "")
	rootDomain := g.Config.Get("dns", "root", "duguying.net")
	rpiRecord := g.Config.Get("dns", "rr", "rpi")
	err = alidns.AddDomainRecord(rootAddr, ak, sk, rootDomain, rpiRecord, "A", 60, "default", c.ClientIP())
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

type AgentDetail struct {
	Id          uint      `json:"id"`
	Online      uint      `json:"online"` // 1 online, 0 offline
	ClientId    string    `json:"client_id" gorm:"unique;not null"`
	Os          string    `json:"os"`
	Arch        string    `json:"arch"`
	Hostname    string    `json:"hostname"`
	Ip          string    `json:"ip"`
	IpIns       []string  `json:"ip_ins"` // json
	Status      uint      `json:"status"`
	OnlineTime  time.Time `json:"online_time"`
	OfflineTime time.Time `json:"offline_time"`
}

func List(c *gin.Context) {
	agents, err := db.ListAllAvailableAgents()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"list": agents,
	})
	return
}
