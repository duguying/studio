// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/5/18.

// Package middleware 中间件
package middleware

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
			sid = c.GetHeader("X-Token")
		}
		// websocket 连接，鉴权从 query 取 token
		if c.GetHeader("Upgrade") == "websocket" {
			sid, _ = c.GetQuery("token")
		}
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
		if entity.UserID <= 0 {
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
			c.Set("user_id", int64(entity.UserID))
		}
		c.Next()
	}
}
