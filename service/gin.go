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
	router.GET("/list", action.ListArticleWithContent)
	router.GET("/list_title", action.ListArticleTitle)
	router.GET("/hot_article", action.HotArticleTitle)
	router.GET("/month_archive", action.MonthArchive)
	router.GET("/user_info", action.UserInfo)

	router.Static("/static/upload", g.Config.Get("upload", "dir", "upload"))

	// print http port
	port := g.Config.GetInt64("system", "port", 9080)
	fmt.Printf("port: %d\n", port)

	router.Run(fmt.Sprintf(":%d", port))
}
