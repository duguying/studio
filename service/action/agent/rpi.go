// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/5/18.

package agent

import (
	"duguying/blog/g"
	"duguying/blog/modules/alidns"
	"github.com/gin-gonic/gin"
	"net/http"
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

	rootDomain := g.Config.Get("dns", "root", "duguying.net")
	rpiRecord := g.Config.Get("dns", "rr", "rpi")
	err = alidns.AddDomainRecord(rootDomain, rpiRecord, "A", 60, "default", c.ClientIP())
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
