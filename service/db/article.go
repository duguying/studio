// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/1/25.

package db

import (
	"duguying/blog/g"
	"duguying/blog/modules/models"
)

func PageArticle(page uint, pageSize uint) (total uint, list []*models.Article, err error) {
	total = 0
	errs := g.Db.Table("articles").Where("status=?", 1).Count(&total).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return 0, nil, errs[0]
	}

	list = []*models.Article{}
	errs = g.Db.Table("articles").Where("status=?", 1).Order("time desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return 0, nil, errs[0]
	}

	return total, list, nil
}

func ArticleToContent(articles []*models.Article) (articleContent []*models.WrapperArticleContent) {
	articleContent = []*models.WrapperArticleContent{}
	for _, article := range articles {
		articleContent = append(articleContent, article.ToArticleContent())
	}
	return articleContent
}

func ArticleToTitle(articles []*models.Article) (articleTitle []*models.WrapperArticleTitle) {
	articleTitle = []*models.WrapperArticleTitle{}
	for _, article := range articles {
		articleTitle = append(articleTitle, article.ToArticleTitle())
	}
	return articleTitle
}