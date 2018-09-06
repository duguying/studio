// Copyright 2017. All rights reserved.
// This file is part of blog project
// Created by duguying on 2017/11/2.

package service

import (
	"duguying/studio/g"
	"duguying/studio/modules/logger"
	"duguying/studio/modules/middleware"
	"duguying/studio/service/action"
	"duguying/studio/service/action/agent"
	"duguying/studio/service/message/deal"
	"duguying/studio/service/message/pipe"
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/gin-contrib/sentry"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func Run() {
	if g.Config.Get("sentry", "enable", "false") == "true" {
		dsn := g.Config.Get("sentry", "dsn", "https://<key>:<secret>@app.getsentry.com/<project>")
		raven.SetDSN(dsn)
	}

	gin.SetMode(g.Config.Get("system", "mode", gin.ReleaseMode))
	gin.DefaultWriter, _ = logger.GinLogger(filepath.Join("log", "gin.log"))

	initWsMessage()

	router := gin.Default()
	router.Use(middleware.CrossSite())
	router.Use(sentry.Recovery(raven.DefaultClient, false))

	router.Any("/version", action.Version)

	api := router.Group("/api")
	{
		api.GET("/get_article", action.GetArticle)
		api.GET("/list", action.ListArticleWithContent)
		api.GET("/list_title", action.ListArticleTitle)
		api.GET("/hot_article", action.HotArticleTitle)
		api.GET("/month_archive", action.MonthArchive)
		api.GET("/user_info", action.UserInfo)
		api.POST("/user_login", action.UserLogin)
		api.Any("/xterm", action.ConnectXTerm)
		api.GET("/username_check", action.UsernameCheck)
		api.POST("/put", action.PutFile)
		api.POST("/upload", action.UploadFile)
		api.GET("/file/list", action.PageFile)

		agt := api.Group("/agent")
		{
			agt.GET("/list", agent.List)
			agt.Any("/ws", agent.Ws)
			agt.GET("/list_perf", agent.PerfList)
		}

		deployer := api.Group("/deploy", action.CheckToken)
		{
			deployer.POST("/upload", action.PackageUpload)
		}
	}

	auth := router.Group("/api/admin", action.SessionValidate)
	{
		auth.POST("/user_logout", action.UserLogout)
	}

	router.Static("/static/upload", g.Config.Get("upload", "dir", "upload"))

	// print http port
	port := g.Config.GetInt64("system", "port", 9080)
	fmt.Printf("port: %d\n", port)

	router.Run(fmt.Sprintf(":%d", port))
}

func initWsMessage() {
	pipe.InitPipeline()
	deal.Start()
	deal.InitHb()
}
