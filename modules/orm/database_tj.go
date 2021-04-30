// Copyright 2017. All rights reserved.
// This file is part of blog project
// Created by duguying on 2017/11/2.

// Package orm ORM初始化包
package orm

import (
	"duguying/studio/g"
	"duguying/studio/modules/dbmodels"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initTjDatabase() {
	if g.Config.SectionExist("database") {
		dbType := g.Config.Get("database", "type", "sqlite")
		if dbType == "mysql" {
			initTjMysql()
		} else {
			initTjSqlite()
		}

		if g.Config.Get("database", "log", "enable") == "enable" {
			g.Db.LogMode(true)
		}

		initTjOrm()
	} else {
		g.InstallMode = true
	}
}

func initTjMysql() {
	var err error
	host := g.Config.Get("database-tj", "host", "127.0.0.1")
	port := g.Config.GetInt64("database-tj", "port", 3306)
	username := g.Config.Get("database-tj", "username", "user")
	password := g.Config.Get("database-tj", "password", "password")
	dbname := g.Config.Get("database-tj", "name", "blog")
	g.Db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname))
	if err != nil {
		log.Printf("数据库连接失败 err:%v\n", err)
	}
}

func initTjSqlite() {
	var err error
	path := g.Config.Get("database-tj", "path", "blog.db")
	g.Db, err = gorm.Open("sqlite3", path)
	if err != nil {
		log.Printf("数据库连接失败 err:%v\n", err)
	}
}

func initTjOrm() {
	g.Db.AutoMigrate(
		&dbmodels.TrojanUsers{},
	)
}
