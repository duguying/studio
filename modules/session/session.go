// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/5/18.

package session

import (
	"duguying/studio/g"
	"duguying/studio/modules/cache"
	"duguying/studio/utils"
	"time"

	"github.com/gogather/json"
)

type Entity struct {
	UserId uint `json:"user_id"`
}

func (se *Entity) String() string {
	c, _ := json.Marshal(se)
	return string(c)
}

func SessionID() string {
	guid, _ := utils.GenUUID()
	return guid
}

func SessionSet(sessionId string, ttl time.Duration, entity *Entity) {
	g.Cache.SetTTL(cache.SESS+sessionId, entity.String(), ttl)
}

func SessionDel(sessionId string) {
	g.Cache.Delete(cache.SESS + sessionId)
}

func SessionGet(sessionId string) (entity *Entity) {
	value, err := g.Cache.Get(cache.SESS + sessionId)
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
