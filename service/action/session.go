// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/5/18.

// Package action 控制器包
package action

import (
	"duguying/studio/g"
	"duguying/studio/modules/session"
	"duguying/studio/service/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SessionValidate(forbidAnonymous bool) func(c *gin.Context) {
	return func(c *gin.Context) {
		sid, err := c.Cookie("sid")
		if err != nil {
			log.Printf("get cookie failed, try to get token, err: %s\n", err.Error())
			sid = c.GetHeader("X-Token")
		}
		log.Printf("get sid: %s\n", sid)
		c.Set("sid", sid)
		sessionDomain := g.Config.Get("session", "domain", ".duguying.net")
		entity := session.SessionGet(sid)
		if entity == nil {
			c.SetCookie("sid", "", 0, "/", sessionDomain, true, false)
			if forbidAnonymous {
				c.JSON(http.StatusUnauthorized, models.CommonResponse{
					Ok:  false,
					Msg: "login first",
				})
				c.Abort()
				return
			} else {
				c.Next()
				return
			}
		} else {
			log.Printf("the entity is: %s\n", entity.String())
		}
		if entity.UserId <= 0 {
			c.SetCookie("sid", "", 0, "/", sessionDomain, true, false)
			session.SessionDel(sid)
			if forbidAnonymous {
				c.JSON(http.StatusUnauthorized, models.CommonResponse{
					Ok:  false,
					Msg: "invalid user",
				})
				c.Abort()
				return
			} else {
				c.Next()
				return
			}
		} else {
			c.Set("user_id", int64(entity.UserId))
		}
		c.Next()
	}
}
