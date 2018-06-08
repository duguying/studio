// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/5/18.

package action

import (
	"duguying/studio/g"
	"duguying/studio/modules/session"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SessionValidate(c *gin.Context) {
	sid, _ := c.Cookie("sid")
	c.Set("sid", sid)
	sessionDomain := g.Config.Get("session", "domain", ".duguying.net")
	entity := session.SessionGet(sid)
	if entity == nil {
		c.SetCookie("sid", "", 0, "/", sessionDomain, true, false)
		c.JSON(http.StatusUnauthorized, gin.H{
			"ok":  false,
			"err": "login first",
		})
		c.Abort()
		return
	}
	if entity.UserId <= 0 {
		c.SetCookie("sid", "", 0, "/", sessionDomain, true, false)
		session.SessionDel(sid)
		c.JSON(http.StatusUnauthorized, gin.H{
			"ok":  false,
			"err": "invalid user",
		})
		c.Abort()
		return
	} else {
		c.Set("user_id", int64(entity.UserId))
	}
	c.Next()
}
