// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/1/25.

package action

import (
	"duguying/studio/modules/db"
	"duguying/studio/modules/dbmodels"
	"duguying/studio/utils"
	"github.com/gin-gonic/gin"
	"github.com/gogather/json"
	"log"
	"net/http"
	"strconv"
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
			"data": art.ToArticleContent(),
		})
		return
	}
}

type ArticleAddRequest struct {
	Title    string
	Keywords []string
	Content  string
	Type     int
	Status   int
}

func (aar *ArticleAddRequest) String() string {
	c, _ := json.Marshal(aar)
	return string(c)
}

func AddArticle(c *gin.Context) {
	aar := &ArticleAddRequest{}
	err := c.BindJSON(aar)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}
	userId := uint(c.GetInt64("user_id"))
	user, err := db.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}
	article, err := db.AddArticle(aar.Title, utils.TitleToUri(aar.Title), aar.Keywords, "", aar.Type, aar.Content, user.Username, user.Id, aar.Status)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"ok":   true,
			"data": article,
		})
		return
	}
}

func PublishArticle(c *gin.Context) {
	aidStr := c.Param("article_id")
	aid64, err := strconv.ParseUint(aidStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}
	aid := uint(aid64)

	// get article
	article, err := db.GetArticleById(aid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	// check article status
	if article.Status != dbmodels.ArtStatus_Publish {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": "it's already published, needn't publish again",
		})
		return
	}

	// check auth
	userId := uint(c.GetInt64("user_id"))
	if userId != article.AuthorId {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": "auth failed, it's not you article, could not publish",
		})
		return
	}

	// publish
	err = db.PublishArticle(aid, userId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"ok":  true,
			"err": "publish success",
		})
		return
	}
}

func DeleteArticle(c *gin.Context) {
	aidStr := c.Param("article_id")
	aid64, err := strconv.ParseUint(aidStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}
	aid := uint(aid64)

	// get article
	article, err := db.GetArticleById(aid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	// check auth
	userId := uint(c.GetInt64("user_id"))
	if userId != article.AuthorId {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": "auth failed, it's not you article, could not publish",
		})
		return
	}

	// delete
	err = db.DeleteArticle(aid, userId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"ok":  true,
			"err": "delete success",
		})
		return
	}
}
