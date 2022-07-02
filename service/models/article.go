// Copyright 2019. All rights reserved.
// This file is part of duguying project
// I am coding in Tencent
// Created by rainesli on 2019/8/29.

package models

import (
	"time"

	"github.com/gogather/json"
)

type Article struct {
	Id       uint     `json:"id"`
	Title    string   `json:"title"`
	Uri      string   `json:"uri"`
	Keywords []string `json:"keywords"`
	Abstract string   `json:"abstract"`
	Content  string   `json:"content"`
	Type     int      `json:"type"`
	Draft    bool     `json:"draft"`
}

func (aar *Article) String() string {
	c, _ := json.Marshal(aar)
	return string(c)
}

type ArticleShowContent struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Uri       string    `json:"uri"`
	Author    string    `json:"author"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	ViewCount uint      `json:"view_count"`
	Content   string    `json:"content"`
}

func (ac *ArticleShowContent) String() string {
	c, _ := json.Marshal(ac)
	return string(c)
}

type ArticleContent struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Uri       string    `json:"uri"`
	Author    string    `json:"author"`
	Tags      []string  `json:"tags"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	ViewCount uint      `json:"view_count"`
	Content   string    `json:"content"`
}

func (asc *ArticleContent) String() string {
	c, _ := json.Marshal(asc)
	return string(c)
}

type ArticleTitle struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Uri       string    `json:"uri"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	ViewCount uint      `json:"view_count"`
}

func (at *ArticleTitle) String() string {
	c, _ := json.Marshal(at)
	return string(c)
}

type ArticleSearchAbstract struct {
	Id        uint       `json:"id"`
	Title     string     `json:"title"`
	Keywords  string     `json:"keywords"`
	Content   string     `json:"content"`
	CreatedAt *time.Time `json:"created_at"`
}

func (asa *ArticleSearchAbstract) String() string {
	c, _ := json.Marshal(asa)
	return string(c)
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
