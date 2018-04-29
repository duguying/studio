// Copyright 2017. All rights reserved.
// This file is part of blog project
// Created by duguying on 2017/11/2.

package orm

import (
	"duguying/blog/g"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"duguying/blog/modules/models"
)

func InitDatabase() {
	if g.Config.SectionExist("database") {
		dbType := g.Config.Get("database", "type", "sqlite")
		if dbType == "mysql" {
			initMysql()
		} else {
			initSqlite()
		}

		initOrm()
	} else {
		g.InstallMode = true
	}
}

func initMysql() {
	var err error
	host := g.Config.Get("database", "host", "127.0.0.1")
	port := g.Config.GetInt64("database", "port", 3306)
	username := g.Config.Get("database", "username", "user")
	password := g.Config.Get("database", "password", "password")
	dbname := g.Config.Get("database", "name", "blog")
	g.Db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname))
	if err != nil {
		log.Printf("数据库连接失败 err:%v\n", err)
	}else {
		if g.Config.Get("database", "log", "enable") == "enable" {
			g.Db.LogMode(true)
		}
	}
}

func initSqlite() {
	var err error
	path := g.Config.Get("database", "path", "blog.db")
	g.Db, err = gorm.Open("sqlite3", path)
	if err != nil {
		log.Printf("数据库连接失败 err:%v\n", err)
	}
}

func initOrm() {
	g.Db.AutoMigrate(&models.Article{},&models.User{},&models.File{})
}