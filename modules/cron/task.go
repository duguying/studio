// Copyright 2020. All rights reserved.
// This file is part of duguying project
// I am coding in Tencent
// Created by rainesli on 2020/3/16.

package cron

import (
	"duguying/studio/g"
	"duguying/studio/modules/db"
	"duguying/studio/modules/viewcnt"
	"duguying/studio/utils"
	"fmt"
	"log"
	"time"

	"github.com/gogather/cron"
)

func Init() {
	task := cron.New()

	spec := g.Config.Get("cron", "flust-view-count", fmt.Sprintf("@every 5m"))
	t1, err := task.AddFunc(spec, flushViewCnt)
	if err != nil {
		log.Println("create cron task failed, err:", err.Error())
	} else {
		log.Println("create cron task success, task id:", t1)
	}

	spec2 := g.Config.Get("calendar", "birth-check", fmt.Sprintf("@daily"))
	t2, err := task.AddFunc(spec2, calendarCheck)
	if err != nil {
		log.Println("create cron task failed, err:", err.Error())
	} else {
		log.Println("create cron task success, task id:", t2)
	}

	task.Start()
}

func flushViewCnt() {
	vcm := viewcnt.GetMap()
	for ident, val := range vcm.M {
		err := db.UpdateArticleViewCount(ident, val.(int))
		if err != nil {
			log.Println("update article view count failed, err:", err.Error())
		} else {
			viewcnt.ResetViewCnt(ident)
		}
	}
}

func calendarCheck() {
	list, err := db.ListAllCalendarIds(g.Db)
	if err != nil {
		log.Println("列举日历事件失败, err:", err.Error())
		return
	}

	beforeDay := g.Config.GetInt64("calendar", "before-day", 7)

	for _, id := range list {
		cal, err := db.GetCalendarById(g.Db, id)
		if err != nil {
			log.Println("获取日历详情失败, err:", err.Error())
			continue
		}
		if cal.Start.Add(-time.Hour * 24 * time.Duration(beforeDay)).Before(time.Now()) {
			utils.GenerateICS(
				cal.Id,
				cal.Start, cal.End, cal.Stamp,
				cal.Summary, cal.Address, cal.Description,
				cal.Link, cal.Attendee,
			)
		}
	}
}
