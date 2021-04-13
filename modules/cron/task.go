// Copyright 2020. All rights reserved.
// This file is part of duguying project
// I am coding in Tencent
// Created by rainesli on 2020/3/16.

package cron

import (
	"duguying/studio/g"
	"duguying/studio/modules/db"
	"duguying/studio/modules/viewcnt"
	"fmt"
	"log"

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

}
