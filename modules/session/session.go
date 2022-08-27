// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/5/18.

// Package session 会话管理
package session

import (
	"duguying/studio/g"
	"duguying/studio/modules/cache"
	"duguying/studio/utils"
	"time"

	"github.com/gogather/json"
)

type Entity struct {
	UserID  uint      `json:"user_id"`
	IP      string    `json:"ip"`
	LoginAt time.Time `json:"login_at"`
}

func (se *Entity) String() string {
	c, _ := json.Marshal(se)
	return string(c)
}

func SessionID() string {
	guid, _ := utils.GenUUID()
	return guid
}

func SessionSet(sessionID string, ttl time.Duration, entity *Entity) {
	g.Cache.SetTTL(cache.SESS+sessionID, entity.String(), ttl)
}

func SessionDel(sessionID string) {
	g.Cache.Delete(cache.SESS + sessionID)
}

func SessionGet(sessionID string) (entity *Entity) {
	value, err := g.Cache.Get(cache.SESS + sessionID)
	if err != nil {
		// log.Println("get session from cache failed, err:", err.Error())
		return nil
	} else {
		entity = &Entity{}
		err = json.Unmarshal([]byte(value), entity)
		if err != nil {
			// log.Println("unmarshal session entity failed, err:", err.Error())
			return nil
		} else {
			return entity
		}
	}
}
