// Copyright 2017. All rights reserved.
// This file is part of blog project
// Created by duguying on 2017/11/2.

package dbmodels

import (
	"duguying/studio/service/models"
	"github.com/gogather/json"
	"gopkg.in/russross/blackfriday.v2"
	"strconv"
	"strings"
	"time"
)

const (
	ArtStatus_Draft   = 0
	ArtStatus_Publish = 1
)

const (
	ContentType_HTML     = 0
	ContentType_MarkDown = 1
)

type Article struct {
	Id          uint       `json:"id"`
	Title       string     `json:"title"`
	Uri         string     `json:"uri"`
	Keywords    string     `json:"keywords"`
	Abstract    string     `json:"abstract"`
	Type        int        `json:"type" gorm:"default:0"`
	Content     string     `json:"content" sql:"type:longtext"`
	Author      string     `json:"author"`
	AuthorId    uint       `json:"author_id"`
	Count       uint       `json:"count"`
	Status      int        `json:"status"`
	PublishTime time.Time  `json:"publish_time"`
	UpdatedBy   uint       `json:"updated_by"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CreatedAt   time.Time  `json:"created_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

func (a *Article) String() string {
	c, _ := json.Marshal(a)
	return string(c)
}

func (a *Article) ToArticleContent() *models.ArticleContent {
	content := []byte(a.Content)
	if a.Type == ContentType_MarkDown {
		content = blackfriday.Run([]byte(a.Content))
	}
	return &models.ArticleContent{
		Id:        a.Id,
		Title:     a.Title,
		Uri:       a.Uri,
		Author:    a.Author,
		Tags:      strings.Split(strings.Replace(a.Keywords, "，", ",", -1), ","),
		CreatedAt: a.CreatedAt,
		ViewCount: a.Count,
		Content:   string(content),
	}
}

func (a *Article) ToArticleTitle() *models.ArticleTitle {
	return &models.ArticleTitle{
		Id:        a.Id,
		Title:     a.Title,
		Uri:       "/article/" + a.Uri,
		Author:    a.Author,
		CreatedAt: a.CreatedAt,
		ViewCount: a.Count,
	}
}

type ArchInfo struct {
	Date   string `json:"date"`
	Number uint   `json:"number"`
	Year   uint   `json:"year"`
	Month  uint   `json:"month"`
}

func (ai *ArchInfo) String() string {
	c, _ := json.Marshal(ai)
	return string(c)
}

func (ai *ArchInfo) parse() {
	segs := strings.Split(ai.Date, "-")
	if len(segs) > 1 {
		month, _ := strconv.ParseInt(segs[1], 10, 32)
		ai.Month = uint(month)
	}
	if len(segs) > 0 {
		year, _ := strconv.ParseInt(segs[0], 10, 32)
		ai.Year = uint(year)
	}
}

func (ai *ArchInfo) ToModel() *models.ArchInfo {
	ai.parse()
	return &models.ArchInfo{
		Date:   ai.Date,
		Number: ai.Number,
		Year:   ai.Year,
		Month:  ai.Month,
	}
}
