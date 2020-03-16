// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/1/25.

package action

import (
	"duguying/studio/g"
	"duguying/studio/modules/db"
	"duguying/studio/modules/viewcnt"
	"duguying/studio/service/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

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

	total, list, err := db.PageArticle("", pager.Page, pager.Size)
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
		Total: total,
		List:  db.ArticleToShowContent(list),
	})
	return
}

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

	total, list, err := db.PageArticle(pager.Tag, pager.Page, pager.Size)
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
		Total: total,
		List:  db.ArticleToShowContent(list),
	})
	return
}

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

	total, list, err := db.PageArticleMonthly(pager.Year, pager.Month, pager.Page, pager.Size)
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
		Total: total,
		List:  db.ArticleToShowContent(list),
	})
	return
}

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

	total, list, err := db.PageArticle("", pager.Page, pager.Size)
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
		Total: total,
		List:  db.ArticleToTitle(list),
	})
	return
}

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

	art := &models.ArticleContent{}
	if getter.Id > 0 {
		dbArt, err := db.GetArticleById(getter.Id)
		if err != nil {
			c.JSON(http.StatusOK, models.ArticleContentGetResponse{
				Ok:  false,
				Msg: err.Error(),
			})
			return
		}
		art = dbArt.ToArticleContent()
	} else if len(getter.Uri) > 0 {
		dbArt, err := db.GetArticle(getter.Uri)
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

	list, err := db.HotArticleTitle(getter.Top)
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

// @Router /month_archive [get]
// @Tags 文章
// @Description 文章按月归档
// @Param top query uint true "前N"
// @Success 200 {object} models.ArticleArchListResponse
func MonthArchive(c *gin.Context) {
	list, err := db.MonthArch()
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

	art := &models.ArticleShowContent{}
	if getter.Id > 0 {
		dbArt, err := db.GetArticleById(getter.Id)
		if err != nil {
			c.JSON(http.StatusOK, models.ArticleShowContentGetResponse{
				Ok:  false,
				Msg: err.Error(),
			})
			return
		}
		art = dbArt.ToArticleShowContent()
	} else if len(getter.Uri) > 0 {
		dbArt, err := db.GetArticle(getter.Uri)
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

// @Router /article/view_count [get]
// @Tags 文章
// @Summary 文章文章浏览统计上报
// @Success 200 {object} models.CommonResponse
func ArticleViewCount(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, models.CommonResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}
	refer := c.GetHeader("Referer")
	referUrl, err := url.Parse(refer)
	if err != nil {
		c.JSON(http.StatusOK, models.CommonResponse{
			Ok:  true,
			Msg: "",
		})
		return
	}
	if referUrl.Host != g.Config.Get("system", "host", "www.duguying.net") {
		c.JSON(http.StatusOK, models.CommonResponse{
			Ok:  true,
			Msg: "",
		})
		return
	}
	viewcnt.ViewHit(int(id))
	c.JSON(http.StatusOK, models.CommonResponse{
		Ok:  true,
		Msg: "",
	})
	return
}

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
	article, err := db.GetArticleById(pub.Id)
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

	// publish
	err = db.PublishArticle(pub.Id, pub.Publish, userId)
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

// @Router /admin/article [delete]
// @Tags 文章
// @Description 删除文章
// @Param id query uint true "文章ID"
// @Success 200 {object} models.CommonResponse
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

	// tag
	tags, counts, err := db.ListAllTags()
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
		for i := uint(1); i <= number; i++ {
			sitemap = append(sitemap, fmt.Sprintf("/tag/%s/%d", tag, number))
		}
	}

	c.JSON(http.StatusOK, models.CommonListResponse{
		Ok:   true,
		List: sitemap,
	})
	return
}
