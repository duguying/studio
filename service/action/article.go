// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/1/25.

package action

import (
	"duguying/studio/modules/db"
	"duguying/studio/modules/dbmodels"
	"duguying/studio/service/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// @Router /list [get]
// @Tags 文章
// @Description 文章列表
// @Param page query uint true "页码"
// @Param size query uint true "每页数"
// @Success 200 {object} models.CommonCreateResponse
func ListArticleWithContent(c *gin.Context) {
	pager := models.CommonPagerRequest{}
	err := c.BindQuery(&pager)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	total, list, err := db.PageArticle(pager.Page, pager.Size)
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

func ListArticleWithContentMonthly(c *gin.Context) {
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

	yearStr := c.Query("year")
	year, err := strconv.ParseUint(yearStr, 10, 64)
	if err != nil {
		log.Println("year 解析错误, err:", err)
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	monthStr := c.Query("month")
	month, err := strconv.ParseUint(monthStr, 10, 64)
	if err != nil {
		log.Println("month 解析错误, err:", err)
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	total, list, err := db.PageArticleMonthly(uint(year), uint(month), uint(page), uint(pageSize))
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

func GetArticleInfo(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	art, err := db.GetArticleById(uint(id))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"item": art,
	})
	return
}

// @Router /admin/article/add [post]
// @Tags 文章
// @Description 创建文章
// @Param area body models.Article true "文章信息"
// @Success 200 {object} models.CommonCreateResponse
func AddArticle(c *gin.Context) {
	aar := &models.Article{}
	err := c.BindJSON(aar)
	if err != nil {
		c.JSON(http.StatusOK, models.CommonCreateResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}

	userId := uint(c.GetInt64("user_id"))
	user, err := db.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusOK, models.CommonCreateResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}

	article, err := db.AddArticle(aar, user.Username, userId)
	if err != nil {
		c.JSON(http.StatusOK, models.CommonCreateResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, models.CommonCreateResponse{
			Ok: true,
			Id: article.Id,
		})
		return
	}
}

func PublishArticle(c *gin.Context) {
	getter := models.CommonGetterRequest{}
	err := c.BindQuery(&getter)
	if err != nil {
		c.JSON(http.StatusOK, models.CommonResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}

	// get article
	article, err := db.GetArticleById(getter.Id)
	if err != nil {
		c.JSON(http.StatusOK, models.CommonResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}

	// check article status
	if article.Status == dbmodels.ArtStatus_Publish {
		c.JSON(http.StatusOK, models.CommonResponse{
			Ok:  false,
			Msg: "it's already published, needn't publish again",
		})
		return
	}

	// check auth
	userId := uint(c.GetInt64("user_id"))
	if userId != article.AuthorId {
		c.JSON(http.StatusOK, models.CommonResponse{
			Ok:  false,
			Msg: "auth failed, it's not you article, could not publish",
		})
		return
	}

	// publish
	err = db.PublishArticle(getter.Id, userId)
	if err != nil {
		c.JSON(http.StatusOK, models.CommonResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, models.CommonResponse{
			Ok:  true,
			Msg: "publish success",
		})
		return
	}
}

func DeleteArticle(c *gin.Context) {
	aidStr := c.Param("article_id")
	aid64, err := strconv.ParseUint(aidStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, models.CommonResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}
	aid := uint(aid64)

	// get article
	article, err := db.GetArticleById(aid)
	if err != nil {
		c.JSON(http.StatusOK, models.CommonResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}

	// check auth
	userId := uint(c.GetInt64("user_id"))
	if userId != article.AuthorId {
		c.JSON(http.StatusOK, models.CommonResponse{
			Ok:  false,
			Msg: "auth failed, it's not you article, could not publish",
		})
		return
	}

	// delete
	err = db.DeleteArticle(aid, userId)
	if err != nil {
		c.JSON(http.StatusOK, models.CommonResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, models.CommonResponse{
			Ok:  true,
			Msg: "delete success",
		})
		return
	}
}

func SiteMap(c *gin.Context) {
	list, err := db.ListAllArticleUri()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	sitemap := []string{
		"/",
		"/about",
	}

	// articles
	for _, item := range list {
		sitemap = append(sitemap, fmt.Sprintf("/article/%s", item.Uri))
	}

	// list pages
	totalPage := len(list) / 40
	if len(list)%40 > 0 {
		totalPage++
	}

	for i := 1; i <= totalPage; i++ {
		sitemap = append(sitemap, fmt.Sprintf("/list/%d", i))
	}

	// pages
	totalArticlePage := len(list) / 10
	if len(list)%10 > 0 {
		totalArticlePage++
	}
	for i := 1; i <= totalArticlePage; i++ {
		sitemap = append(sitemap, fmt.Sprintf("/page/%d", i))
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"list": sitemap,
	})
	return
}
