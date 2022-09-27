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
	"time"

	"github.com/gogather/d2"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	cache = d2.NewD2()
)

func InitDatabase() {
	initDatabase()
}

func initDatabase() {
	if g.Config.SectionExist("database") {
		dbType := g.Config.Get("database", "type", "sqlite")
		if dbType == "mysql" {
			initMysql()
		} else {
			initSqlite()
		}

		if g.Config.Get("database", "log", "enable") == "enable" {
			// g.Db.LogMode(true)
		}

		initOrm()
	} else {
		g.InstallMode = true
	}
}

func initMysql() {
	newLogger := New(
		Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // Disable color
		},
	)

	var err error
	host := g.Config.Get("database", "host", "127.0.0.1")
	port := g.Config.GetInt64("database", "port", 3306)
	username := g.Config.Get("database", "username", "user")
	password := g.Config.Get("database", "password", "password")
	dbname := g.Config.Get("database", "name", "blog")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)

	g.Db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256, // default size for string fields
	}), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Fatalf("数据库连接失败 err:%v\n", err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db, _ := g.Db.DB()
	db.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.SetConnMaxLifetime(time.Hour)
}

func initSqlite() {
	var err error
	path := g.Config.Get("database", "path", "blog.db")
	g.Db, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		log.Printf("数据库连接失败 err:%v\n", err)
	}
}

func initOrm() {
	g.Db.AutoMigrate(
		&dbmodels.Article{},
		&dbmodels.Draft{},
		&dbmodels.User{},
		&dbmodels.File{},
		&dbmodels.Agent{},
		&dbmodels.AgentPerform{},
		&dbmodels.Face{},
		&dbmodels.FaceLabel{},
		&dbmodels.LoginHistory{},
		&dbmodels.ImageMeta{},
		&dbmodels.Cover{},
		&dbmodels.Calendar{},
	)
}
