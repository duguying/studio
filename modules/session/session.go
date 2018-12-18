// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/5/18.

package session

import (
	"duguying/studio/modules/redis"
	"duguying/studio/utils"
	"github.com/gogather/json"
	"log"
	"time"
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
	redis.SetTTL(redis.SESS+sessionId, entity.String(), ttl)
}

func SessionDel(sessionId string) {
	redis.Delete(redis.SESS + sessionId)
}

func SessionGet(sessionId string) (entity *Entity) {
	value, err := redis.Get(redis.SESS + sessionId)
	if err != nil {
		log.Println("get session from redis failed, err:", err.Error())
		return nil
	} else {
		entity = &Entity{}
		err = json.Unmarshal([]byte(value), entity)
		if err != nil {
			log.Println("unmarshal session entity failed, err:", err.Error())
			return nil
		} else {
			return entity
		}
	}
}
