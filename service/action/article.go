// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/1/25.

package action

import (
	"duguying/studio/g"
	"duguying/studio/modules/db"
	"duguying/studio/modules/viewcnt"
	"duguying/studio/service/models"
	"duguying/studio/utils"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
)

// ListArticleWithContent 文章列表
// @Router /list [get]
// @Tags 文章
// @Description 文章列表
// @Param page query uint true "页码"
// @Param size query uint true "每页数"
// @Success 200 {object} models.ArticleContentListResponse
func ListArticleWithContent(c *gin.Context) {
	pager := models.CommonPagerRequest{}
	err := c.BindQuery(&pager)
	if err != nil {
		c.JSON(http.StatusOK, models.ArticleContentListResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}

	total, list, err := db.PageArticle(g.Db, "", pager.Page, pager.Size)
	if err != nil {
		log.Println("分页查询错误, err:", err)
		c.JSON(http.StatusOK, models.ArticleContentListResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ArticleContentListResponse{
		Ok:    true,
		Total: uint(total),
		List:  db.ArticleToShowContent(list),
	})
	return
}

// SearchArticle 文章搜索
// @Router /search_article [get]
// @Tags 文章
// @Description 文章搜索
// @Param keyword query string true "关键词"
// @Param page query uint true "页码"
// @Param size query uint true "每页数"
// @Success 200 {object} models.ArticleContentListResponse
func SearchArticle(c *gin.Context) {
	req := models.SearchPagerRequest{}
	err := c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, models.CommonResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}

	// 关键词为空，直接返回空
	if len(req.Keyword) <= 0 {
		c.JSON(http.StatusOK, models.CommonSearchListResponse{
			Ok: true,
		})
		return
	}

	// 默认参数
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 10
	}

	tx := g.Db.WithContext(c)
	total, result, articleMap, err := db.SearchArticle(tx, req.Keyword, req.Page, req.Size)
	if err != nil {
		c.JSON(http.StatusOK, models.CommonResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}

	searchList := []*models.ArticleSearchAbstract{}
	for _, item := range result.Hits {
		id, err := strconv.ParseUint(item.ID, 10, 64)
		if err != nil {
			continue
		}
		article := articleMap[uint(id)]
		title := article.Title
		if len(item.Fragments["title"]) > 0 {
			title = item.Fragments["title"][0]
		}
		keywords := article.Keywords
		if len(item.Fragments["keywords"]) > 0 {
			keywords = item.Fragments["keywords"][0]
		}
		content := utils.TrimHtml(article.Content)
		if utf8.RuneCountInString(content) > 100 {
			content = string([]rune(content)[:100])
		}
		if len(item.Fragments["content"]) > 0 {
			content = item.Fragments["content"][0]
		}
		article.Keywords = keywords // 包含mark标签
		searchList = append(searchList, &models.ArticleSearchAbstract{
			ID:        uint(id),
			Title:     title,
			URI:       article.URI,
			Tags:      article.ToArticleContent().Tags,
			Author:    article.Author,
			Keywords:  keywords,
			Content:   content,
			CreatedAt: &article.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, models.CommonSearchListResponse{
		Ok:    true,
		Total: total,
		List:  searchList,
	})
	return
}

// ListArticleWithContentByTag 通过tag列举文章
// @Router /list_tag [get]
// @Tags 文章
// @Description 文章列表
// @Param tag query string true "Tag"
// @Param page query uint true "页码"
// @Param size query uint true "每页数"
// @Success 200 {object} models.ArticleContentListResponse
func ListArticleWithContentByTag(c *gin.Context) {
	pager := models.TagPagerRequest{}
	err := c.BindQuery(&pager)
	if err != nil {
		c.JSON(http.StatusOK, models.ArticleContentListResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}

	total, list, err := db.PageArticle(g.Db, pager.Tag, pager.Page, pager.Size)
	if err != nil {
		log.Println("分页查询错误, err:", err)
		c.JSON(http.StatusOK, models.ArticleContentListResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ArticleContentListResponse{
		Ok:    true,
		Total: uint(total),
		List:  db.ArticleToShowContent(list),
	})
	return
}

// ListArticleWithContentMonthly 文章按月列表
// @Router /list_archive_monthly [get]
// @Tags 文章
// @Description 文章列表
// @Param page query uint true "页码"
// @Param size query uint true "每页数"
// @Param year query uint true "年"
// @Param month query uint true "月"
// @Success 200 {object} models.ArticleContentListResponse
func ListArticleWithContentMonthly(c *gin.Context) {
	pager := models.MonthlyPagerRequest{}
	err := c.BindQuery(&pager)
	if err != nil {
		c.JSON(http.StatusOK, models.ArticleContentListResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}

	total, list, err := db.PageArticleMonthly(g.Db, pager.Year, pager.Month, pager.Page, pager.Size)
	if err != nil {
		log.Println("分页查询错误, err:", err)
		c.JSON(http.StatusOK, models.ArticleContentListResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ArticleContentListResponse{
		Ok:    true,
		Total: uint(total),
		List:  db.ArticleToShowContent(list),
	})
	return
}

// ListArticleTitle 按标题列举文章
// @Router /list_title [get]
// @Tags 文章
// @Description 文章列表
// @Param page query uint true "页码"
// @Param size query uint true "每页数"
// @Success 200 {object} models.ArticleTitleListResponse
func ListArticleTitle(c *gin.Context) {
	pager := models.CommonPagerRequest{}
	err := c.BindQuery(&pager)
	if err != nil {
		c.JSON(http.StatusOK, models.ArticleTitleListResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}

	total, list, err := db.PageArticle(g.Db, "", pager.Page, pager.Size)
	if err != nil {
		log.Println("分页查询错误, err:", err)
		c.JSON(http.StatusOK, models.ArticleTitleListResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ArticleTitleListResponse{
		Ok:    true,
		Total: uint(total),
		List:  db.ArticleToTitle(list),
	})
	return
}

// GetArticle 获取文章
// @Router /article [get]
// @Tags 文章
// @Description 获取文章
// @Param id query uint false "ID"
// @Param uri query string false "URI"
// @Success 200 {object} models.ArticleContentGetResponse
func GetArticle(c *gin.Context) {
	getter := models.ArticleUriGetterRequest{}
	err := c.BindQuery(&getter)
	if err != nil {
		c.JSON(http.StatusOK, models.ArticleContentGetResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}

	var art *models.ArticleContent
	if getter.Id > 0 {
		dbArt, err := db.GetArticleById(g.Db, getter.Id)
		if err != nil {
			c.JSON(http.StatusOK, models.ArticleContentGetResponse{
				Ok:  false,
				Msg: err.Error(),
			})
			return
		}
		art = dbArt.ToArticleContent()
	} else if len(getter.Uri) > 0 {
		dbArt, err := db.GetArticle(g.Db, getter.Uri)
		if err != nil {
			c.JSON(http.StatusOK, models.ArticleContentGetResponse{
				Ok:  false,
				Msg: err.Error(),
			})
			return
		}
		art = dbArt.ToArticleContent()
	} else {
		c.JSON(http.StatusOK, models.ArticleContentGetResponse{
			Ok:  false,
			Msg: "invalid id and uri",
		})
		return
	}

	c.JSON(http.StatusOK, models.ArticleContentGetResponse{
		Ok:   true,
		Data: art,
	})
	return
}

// HotArticleTitle 文章TopN列表
// @Router /hot_article [get]
// @Tags 文章
// @Description 文章TopN列表
// @Param top query uint true "前N"
// @Success 200 {object} models.ArticleTitleListResponse
func HotArticleTitle(c *gin.Context) {
	getter := models.TopGetterRequest{}
	err := c.BindQuery(&getter)
	if err != nil {
		c.JSON(http.StatusOK, models.ArticleTitleListResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}

	list, err := db.HotArticleTitle(g.Db, getter.Top)
	if err != nil {
		c.JSON(http.StatusOK, models.ArticleTitleListResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.ArticleTitleListResponse{
		Ok:   true,
		List: list,
	})
	return
}

// MonthArchive 文章按月归档
// @Router /month_archive [get]
// @Tags 文章
// @Description 文章按月归档
// @Param top query uint true "前N"
// @Success 200 {object} models.ArticleArchListResponse
func MonthArchive(c *gin.Context) {
	list, err := db.MonthArch(g.Db)
	if err != nil {
		c.JSON(http.StatusOK, models.ArticleArchListResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}

	apiList := []*models.ArchInfo{}
	for _, item := range list {
		apiList = append(apiList, item.ToModel())
	}

	c.JSON(http.StatusOK, models.ArticleArchListResponse{
		Ok:   true,
		List: apiList,
	})
	return
}

// GetArticleShow 获取文章
// @Router /get_article [get]
// @Tags 文章
// @Description 获取文章
// @Param id query uint false "ID"
// @Param uri query string false "URI"
// @Success 200 {object} models.ArticleShowContentGetResponse
func GetArticleShow(c *gin.Context) {
	getter := models.ArticleUriGetterRequest{}
	err := c.BindQuery(&getter)
	if err != nil {
		c.JSON(http.StatusOK, models.ArticleShowContentGetResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}

	tx := g.Db.WithContext(c)
	var art *models.ArticleShowContent
	if getter.Id > 0 {
		dbArt, err := db.GetArticleById(tx, getter.Id)
		if err != nil {
			c.JSON(http.StatusOK, models.ArticleShowContentGetResponse{
				Ok:  false,
				Msg: err.Error(),
			})
			return
		}
		art = dbArt.ToArticleShowContent()
	} else if len(getter.Uri) > 0 {
		dbArt, err := db.GetArticle(tx, getter.Uri)
		if err != nil {
			c.JSON(http.StatusOK, models.ArticleShowContentGetResponse{
				Ok:  false,
				Msg: err.Error(),
			})
			return
		}
		art = dbArt.ToArticleShowContent()
	} else {
		c.JSON(http.StatusOK, models.ArticleShowContentGetResponse{
			Ok:  false,
			Msg: "invalid id and uri",
		})
		return
	}

	c.JSON(http.StatusOK, models.ArticleShowContentGetResponse{
		Ok:   true,
		Data: art,
	})
	return
}

// ArticleViewCount 文章文章浏览统计上报
// @Router /article/view_count [get]
// @Tags 文章
// @Summary 文章文章浏览统计上报
// @Success 200 {object} models.CommonResponse
func ArticleViewCount(c *gin.Context) {
	ident := c.Query("ident")
	refer := c.GetHeader("Referer")
	referURL, err := url.Parse(refer)
	if err != nil {
		c.JSON(http.StatusOK, models.CommonResponse{
			Ok:  true,
			Msg: "",
		})
		return
	}
	if referURL.Host != g.Config.Get("system", "host", "www.duguying.net") {
		c.JSON(http.StatusOK, models.CommonResponse{
			Ok:  true,
			Msg: "",
		})
		return
	}
	viewcnt.ViewHit(ident)
	c.JSON(http.StatusOK, models.CommonResponse{
		Ok:  true,
		Msg: "",
	})
	return
}

// AddArticle 创建文章
// @Router /admin/article [post]
// @Tags 文章
// @Description 创建文章
// @Param article body models.Article true "文章信息"
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

	tx := g.Db.WithContext(c)
	userID := uint(c.GetInt64("user_id"))
	user, err := db.GetUserById(tx, userID)
	if err != nil {
		c.JSON(http.StatusOK, models.CommonCreateResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}

	article, err := db.AddArticle(tx, aar, user.Username, userID)
	if err != nil {
		c.JSON(http.StatusOK, models.CommonCreateResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, models.CommonCreateResponse{
			Ok: true,
			Id: article.ID,
		})
		return
	}
}

// UpdateArticle 修改文章
// @Router /admin/article [put]
// @Tags 文章
// @Description 修改文章
// @Param publish body models.Article true "文章信息"
// @Success 200 {object} models.CommonResponse
func UpdateArticle(c *gin.Context) {
	article := models.Article{}
	err := c.BindJSON(&article)
	if err != nil {
		c.JSON(http.StatusOK, models.CommonResponse{
			Ok:  false,
			Msg: "解析参数失败",
		})
		return
	}
	tx := g.Db.WithContext(c)
	err = db.UpdateArticle(tx, article.ID, &article)
	if err != nil {
		c.JSON(http.StatusOK, models.CommonResponse{
			Ok:  false,
			Msg: "解析参数失败",
		})
		return
	}
	c.JSON(http.StatusOK, models.CommonResponse{
		Ok: true,
	})
	return
}

// PublishArticle 发布文章
// @Router /admin/article/publish [put]
// @Tags 文章
// @Description 发布文章
// @Param publish body models.ArticlePublishRequest true "文章信息"
// @Success 200 {object} models.CommonResponse
func PublishArticle(c *gin.Context) {
	pub := models.ArticlePublishRequest{}
	err := c.BindJSON(&pub)
	if err != nil {
		c.JSON(http.StatusOK, models.CommonResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}

	// get article
	tx := g.Db.WithContext(c)
	article, err := db.GetArticleById(tx, pub.Id)
	if err != nil {
		c.JSON(http.StatusOK, models.CommonResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}

	// check auth
	userID := uint(c.GetInt64("user_id"))
	if userID != article.AuthorID {
		c.JSON(http.StatusOK, models.CommonResponse{
			Ok:  false,
			Msg: "auth failed, it's not you article, could not publish",
		})
		return
	}

	// publish
	err = db.PublishArticle(tx, pub.Id, pub.Publish, userID)
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

// DeleteArticle 删除文章
// @Router /admin/article [delete]
// @Tags 文章
// @Description 删除文章
// @Param id query uint true "文章ID"
// @Success 200 {object} models.CommonResponse
func DeleteArticle(c *gin.Context) {
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
	tx := g.Db.WithContext(c)
	article, err := db.GetArticleById(tx, getter.Id)
	if err != nil {
		c.JSON(http.StatusOK, models.CommonResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}

	// check auth
	userID := uint(c.GetInt64("user_id"))
	if userID != article.AuthorID {
		c.JSON(http.StatusOK, models.CommonResponse{
			Ok:  false,
			Msg: "auth failed, it's not you article, could not publish",
		})
		return
	}

	// delete
	err = db.DeleteArticle(tx, getter.Id, userID)
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

// SiteMap 站点地图
func SiteMap(c *gin.Context) {
	tx := g.Db.WithContext(c)
	list, err := db.ListAllArticleUri(tx)
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
		sitemap = append(sitemap, fmt.Sprintf("/article/%s", item.URI))
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

	// tag
	tags, counts, err := db.ListAllTags(tx)
	if err != nil {
		c.JSON(http.StatusOK, models.CommonListResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}
	for idx, tag := range tags {
		total := counts[idx]
		number := total / 10
		if total%10 > 0 {
			number++
		}
		for i := uint(1); i <= uint(number); i++ {
			sitemap = append(sitemap, fmt.Sprintf("/tag/%s/%d", tag, number))
		}
	}

	c.JSON(http.StatusOK, models.CommonListResponse{
		Ok:   true,
		List: sitemap,
	})
	return
}
