// Copyright 2017. All rights reserved.
// This file is part of blog project
// Created by duguying on 2017/11/2.

package service

import (
	"duguying/blog/g"
	"duguying/blog/modules/logger"
	"duguying/blog/service/action"
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"duguying/blog/modules/middleware"
)

func Run() {
	gin.SetMode(g.Config.Get("system", "mode", gin.ReleaseMode))
	gin.DefaultWriter, _ = logger.GinLogger(filepath.Join("log", "gin.log"))

	router := gin.Default()
	router.Use(middleware.CrossSite())

	router.Any("/version", action.Version)

	api:=router.Group("/api")
	{
		api.GET("/list", action.ListArticleWithContent)
		api.GET("/list_title", action.ListArticleTitle)
		api.GET("/hot_article", action.HotArticleTitle)
		api.GET("/month_archive", action.MonthArchive)
		api.GET("/user_info", action.UserInfo)
		api.POST("/user_login", action.UserLogin)

		rpi:=api.Group("/rpi")
		{
			rpi.POST("/agent", )
		}
	}

	auth:=router.Group("/api/admin", action.SessionValidate)
	{
		auth.POST("/user_logout", action.UserLogout)
	}

	router.Static("/static/upload", g.Config.Get("upload", "dir", "upload"))

	// print http port
	port := g.Config.GetInt64("system", "port", 9080)
	fmt.Printf("port: %d\n", port)

	router.Run(fmt.Sprintf(":%d", port))
}
