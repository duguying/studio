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
	"github.com/gogather/cron"
	"log"
	"strconv"
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

	task.Start()
}

func flushViewCnt() {
	vcm := viewcnt.GetMap()
	for idStr, val := range vcm.M {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			log.Println("parse article id failed, err:", err.Error())
			continue
		} else {
			err = db.UpdateArticleViewCount(uint(id), val.(int))
			if err != nil {
				log.Println("update article view count failed, err:", err.Error())
			}
		}
	}
}
