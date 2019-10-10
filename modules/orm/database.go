// Copyright 2017. All rights reserved.
// This file is part of blog project
// Created by duguying on 2017/11/2.

package orm

import (
	"duguying/studio/g"
	"duguying/studio/modules/dbmodels"
	"encoding/json"
	"fmt"
	"github.com/gogather/d2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var (
	cache = d2.NewD2()
)

func InitDatabase() {
	if g.Config.SectionExist("database") {
		dbType := g.Config.Get("database", "type", "sqlite")
		if dbType == "mysql" {
			initMysql()
		} else {
			initSqlite()
		}

		if g.Config.Get("database", "log", "enable") == "enable" {
			g.Db.LogMode(true)
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
	g.Db.AutoMigrate(
		&dbmodels.Article{},
		&dbmodels.User{},
		&dbmodels.File{},
		&dbmodels.Agent{},
		&dbmodels.AgentPerform{},
		&dbmodels.ApiLog{},
	)

	// clear cache
	g.Db.Callback().Create().After("gorm:create").Register("plugin:run_after_create", func(scope *gorm.Scope) {
		cache.RemoveSection(scope.TableName())
	})

	g.Db.Callback().Update().After("gorm:update").Register("plugin:run_after_update", func(scope *gorm.Scope) {
		cache.RemoveSection(scope.TableName())
	})

	g.Db.Callback().Delete().After("gorm:delete").Register("plugin:run_after_delete", func(scope *gorm.Scope) {
		cache.RemoveSection(scope.TableName())
	})

	// todo: query cache
	g.Db.Callback().Query().Before("gorm:query").Register("plugin:run_before_query", func(scope *gorm.Scope) {
		scope.CallMethod("prepareQuerySQL")
		//fmt.Println("bef SQL:", scope.SQL, "scope:", scope)
		cacheRaw, exist := cache.Get(scope.TableName(), scope.SQL)
		if exist {
			err := json.Unmarshal([]byte(cacheRaw.(string)), scope.Value)
			if err != nil {
				fmt.Println("err:", err.Error())
			}
			scope.Log("Hit Cache", scope)
			scope.SkipLeft()
		}

	})

	// set cache
	g.Db.Callback().Query().After("gorm:query").Register("plugin:run_after_query", func(scope *gorm.Scope) {
		c, _ := json.Marshal(scope.Value)
		//fmt.Println("SQL:", scope.SQL)
		cache.Add(scope.TableName(), scope.SQL, string(c))
	})
}
