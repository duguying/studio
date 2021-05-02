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

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initTjDatabase() {
	if g.Config.SectionExist("database") {
		dbType := g.Config.Get("database-tj", "type", "sqlite")
		if dbType == "mysql" {
			initTjMysql()
		} else {
			initTjSqlite()
		}

		if g.Config.Get("database-tj", "log", "enable") == "enable" {
			// g.Db.LogMode(true)
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
	dbname := g.Config.Get("database-tj", "name", "tgw")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)
	g.GfwDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("数据库连接失败 err:%v\n", err)
	}
}

func initTjSqlite() {
	var err error
	path := g.Config.Get("database-tj", "path", "gfw.db")
	g.GfwDb, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		log.Printf("数据库连接失败 err:%v\n", err)
	}
}

func initTjOrm() {
	g.GfwDb.AutoMigrate(
		&dbmodels.TrojanUsers{},
	)
}
