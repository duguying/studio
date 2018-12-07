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
	"github.com/gin-contrib/pprof"
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
	router.Use(middleware.ServerMark())
	router.Use(middleware.CrossSite())
	router.Use(sentry.Recovery(raven.DefaultClient, false))

	router.Any("/version", action.Version)

	// v1 api
	apiV1 := router.Group("/api/v1")
	{
		// needn't auth
		{
			apiV1.GET("/get_article", action.GetArticle)       // 获取文章详情
			apiV1.GET("/list", action.ListArticleWithContent)  // 列出文章
			apiV1.GET("/list_title", action.ListArticleTitle)  // 列出文章标题
			apiV1.GET("/hot_article", action.HotArticleTitle)  // 热门文章列表
			apiV1.GET("/month_archive", action.MonthArchive)   // 文章按月归档列表
			apiV1.POST("/user_register", action.UserRegister)  // 用户注册
			apiV1.GET("/user_info", action.UserInfo)           // 用户信息
			apiV1.POST("/user_login", action.UserLogin)        // 用户登陆
			apiV1.GET("/username_check", action.UsernameCheck) // 用户名检查
			apiV1.GET("/file/list", action.PageFile)           // 文件列表
			apiV1.POST("/2fa", action.TfaAuth)                 // 2FA校验
		}

		// auth require
		auth := apiV1.Group("", action.SessionValidate)
		{
			auth.POST("/user_logout", action.UserLogout) // 用户登出
			apiV1.POST("/put", action.PutFile)           // 上传文件
			apiV1.POST("/upload", action.UploadFile)     // 上传文件
			apiV1.Any("/xterm", action.ConnectXTerm)     // 连接xterm

			apiV1.POST("/article/add", action.AddArticle)         // 添加文章
			apiV1.POST("/article/publish", action.PublishArticle) // 发布草稿
			apiV1.POST("/article/delete", action.DeleteArticle)   // 删除文章

			apiV1.GET("/2faqr", action.QrGoogleAuth) // 获取2FA二维码
		}

		// agent connection point
		agt := apiV1.Group("/agent")
		{
			agt.GET("/list", action.SessionValidate, agent.List)          // agent列表
			agt.GET("/list_perf", action.SessionValidate, agent.PerfList) // 性能列表
			agt.Any("/ws", agent.Ws)                                      // agent连接点
		}

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
	port := g.Config.GetInt64("system", "port", 9080)
	fmt.Printf("port: %d\n", port)

	pprof.Register(router)
	err := router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("run gin server failed, err:" + err.Error())
	}
}

func initWsMessage() {
	pipe.InitPipeline()
	deal.Start()
	deal.InitHb()
}
