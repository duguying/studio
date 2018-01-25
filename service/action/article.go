// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/1/25.

package action

import (
	"github.com/gin-gonic/gin"
	"duguying/blog/service/db"
	"strconv"
	"net/http"
	"log"
)

func ListArticleWithContent(c *gin.Context) {
	pageStr:=c.Query("page")
	page,err:=strconv.ParseUint(pageStr,10,64)
	if err != nil {
		log.Println("page 解析错误, err:", err)
		c.JSON(http.StatusOK, gin.H{
			"ok":false,
			"err": err.Error(),
		})
		return
	}

	pageSizeStr:=c.Query("page_size")
	pageSize,err:=strconv.ParseUint(pageSizeStr,10,64)
	if err != nil {
		log.Println("page_size 解析错误, err:", err)
		c.JSON(http.StatusOK, gin.H{
			"ok":false,
			"err": err.Error(),
		})
		return
	}

	total,list,err:= db.PageArticle(uint(page), uint(pageSize))
	if err != nil {
		log.Println("分页查询错误, err:", err)
		c.JSON(http.StatusOK, gin.H{
			"ok":false,
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":true,
		"total": total,
		"list": list,
	})
	return
}
