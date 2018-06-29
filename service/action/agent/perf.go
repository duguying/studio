// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/6/8.

package agent

import (
	"duguying/studio/service/message/model"
	"duguying/studio/service/message/store"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"net/http"
)

func PerfList(c *gin.Context) {
	clientId := c.Query("client_id")
	list, err := store.ListPerf(clientId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	} else {
		perfs := []*model.PerformanceMonitor{}
		for _, item := range list {
			perf := model.PerformanceMonitor{}
			err := proto.Unmarshal(item, &perf)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"ok":  false,
					"err": err.Error(),
				})
				return
			}
			perfs = append(perfs, &perf)
		}

		c.JSON(http.StatusOK, gin.H{
			"ok":   true,
			"list": perfs,
		})
		return
	}
}
