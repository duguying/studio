// Copyright 2017. All rights reserved.
// This file is part of blog project
// Created by duguying on 2017/11/2.

// Package service 服务包
package service

import (
	_ "duguying/studio/docs"
	"duguying/studio/g"
	"duguying/studio/modules/logger"
	"duguying/studio/service/action"
	"duguying/studio/service/action/agent"
	"duguying/studio/service/message/deal"
	"duguying/studio/service/message/pipe"
	"duguying/studio/service/middleware"
	"duguying/studio/service/models"
	"fmt"
	"path/filepath"

	"github.com/getsentry/raven-go"
	"github.com/gin-contrib/pprof"
	"github.com/gin-contrib/sentry"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Run(logDir string) {
	models.RegisterTimeAsLayoutCodec("2006-01-02 15:04:05")
	gin.SetMode(g.Config.Get("system", "mode", gin.ReleaseMode))
	gin.DefaultWriter, _ = logger.GinLogger(filepath.Join(logDir, "gin.log"))

	initWsMessage()

	router := gin.Default()
	router.Use(middleware.ServerMark())
	router.Use(middleware.CrossSite())
	router.Use(sentry.Recovery(raven.DefaultClient, false))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Any("/version", action.Version)

	// v1 api
	apiV1 := router.Group("/api/v1", middleware.RestLog(), middleware.SessionValidate(false))
	{
		// needn't auth
		action.SetupFeAPI(apiV1)

		// auth require
		auth := apiV1.Group("/admin", middleware.RestLog(), middleware.SessionValidate(true))
		action.SetupAdminAPI(auth)

		// agent connection point
		agt := apiV1.Group("/agent")
		action.SetupAgentAPI(agt)
	}

	// 兼容旧版
	api := router.Group("/api")
	{
		// 旧版 agent 连接点
		agt := api.Group("/agent")
		{
			agt.Any("/ws", agent.Ws)
		}

		// 静态站点部署器
		deployer := api.Group("/deploy", action.CheckToken)
		{
			deployer.POST("/upload", action.PackageUpload)
			deployer.POST("/archive", action.UploadFile)
		}
	}

	router.Static("/static/upload", g.Config.Get("upload", "dir", "upload"))

	// print http port
	addr := g.Config.Get("system", "listen", "127.0.0.1:9080")
	fmt.Printf("listen: %s\n", addr)

	pprof.Register(router)
	err := router.Run(addr)
	if err != nil {
		fmt.Println("run gin server failed, err:" + err.Error())
	}
}

func initWsMessage() {
	pipe.InitPipeline()
	deal.Start()
	deal.InitHb()
}
