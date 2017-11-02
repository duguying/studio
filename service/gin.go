// Copyright 2017. All rights reserved.
// This file is part of blog project
// Created by duguying on 2017/11/2.

package service

import (
	"github.com/gin-gonic/gin"
	"duguying/blog/g"
	"duguying/blog/modules/logger"
	"duguying/blog/service/action"
	"path/filepath"
	"fmt"
)

func Run() {
	gin.SetMode(g.Config.Get("system", "mode", gin.ReleaseMode))
	gin.DefaultWriter, _ = logger.GinLogger(filepath.Join("log", "gin.log"))

	router := gin.Default()

	router.Any("/version", action.Version)

	router.Static("/static/avatar",g.Config.Get("avatar","dir", "avatar"))

	// print http port
	port := g.Config.GetInt64("system", "port", 9080)
	fmt.Printf("port: %d\n", port)

	router.Run(fmt.Sprintf(":%d", port))
}

