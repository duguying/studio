// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/1/25.

package action

import (
	"duguying/blog/service/db"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"fmt"
)

func ListArticleWithContent(c *gin.Context) {
	pageStr := c.Query("page")
	page, err := strconv.ParseUint(pageStr, 10, 64)
	if err != nil {
		log.Println("page 解析错误, err:", err)
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	pageSizeStr := c.Query("page_size")
	pageSize, err := strconv.ParseUint(pageSizeStr, 10, 64)
	if err != nil {
		log.Println("page_size 解析错误, err:", err)
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	total, list, err := db.PageArticle(uint(page), uint(pageSize))
	if err != nil {
		log.Println("分页查询错误, err:", err)
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":    true,
		"total": total,
		"list":  db.ArticleToContent(list),
	})
	return
}

func ListArticleTitle(c *gin.Context) {
	pageStr := c.Query("page")
	page, err := strconv.ParseUint(pageStr, 10, 64)
	if err != nil {
		log.Println("page 解析错误, err:", err)
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	pageSizeStr := c.Query("page_size")
	pageSize, err := strconv.ParseUint(pageSizeStr, 10, 64)
	if err != nil {
		log.Println("page_size 解析错误, err:", err)
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	total, list, err := db.PageArticle(uint(page), uint(pageSize))
	if err != nil {
		log.Println("分页查询错误, err:", err)
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":    true,
		"total": total,
		"list":  db.ArticleToTitle(list),
	})
	return
}

func HotArticleTitle(c *gin.Context) {
	topStr := c.DefaultQuery("top", "10")
	top, err := strconv.ParseUint(topStr, 10, 64)
	if err != nil {
		log.Printf("解析错误, err: %s\n", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}
	list, err := db.HotArticleTitle(uint(top))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"list": list,
	})
	return
}

func MonthArchive(c *gin.Context) {
	list, err := db.MonthArch()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}
	log.Println("按月归档 list:", list)
	c.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"list": list,
	})
	return
}

func GetArticle(c *gin.Context) {
	uri := c.Query("uri")
	fmt.Println("[URL]", c.Request.URL.Query())
	art, err := db.GetArticle(uri)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"ok":   true,
			"data": art,
		})
		return
	}
}
