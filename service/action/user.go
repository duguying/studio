// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/1/25.

package action

import (
	"duguying/studio/g"
	"duguying/studio/modules/session"
	"duguying/studio/service/db"
	"github.com/gin-gonic/gin"
	"github.com/gogather/com"
	"net/http"
	"time"
)

func UserInfo(c *gin.Context) {

}

type LoginArgs struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func UserLogin(c *gin.Context) {
	login := &LoginArgs{}
	err := c.BindJSON(login)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}
	user, err := db.GetUser(login.Username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	// validate
	passwd := com.Md5(login.Password + user.Salt)
	if passwd != user.Password {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": "login failed, invalid password",
		})
		return
	} else {
		sid := session.SessionID()
		if sid == "" {
			c.JSON(http.StatusOK, gin.H{
				"ok":  false,
				"err": "generate sid failed",
			})
			return
		} else {
			defaultSessionTime := time.Hour * 24
			sessionTimeCfg := g.Config.Get("session", "expire", defaultSessionTime.String())
			sessionExpire, err := time.ParseDuration(sessionTimeCfg)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"ok":  false,
					"err": err.Error(),
				})
				return
			} else {
				// store session
				session.SessionSet(sid, sessionExpire, &session.Entity{
					UserId: user.Id,
				})

				// store cookie
				sessionDomain := g.Config.Get("session", "domain", ".duguying.net")
				c.SetCookie("sid", sid, int(sessionExpire.Seconds()), "/", sessionDomain, true, false)

				c.JSON(http.StatusOK, gin.H{
					"ok": true,
				})
				return
			}

		}
	}
}

func UserLogout(c *gin.Context) {
	sid := c.GetString("sid")
	userId := uint(c.GetInt64("user_id"))
	session.SessionDel(sid)
	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "logout success",
		"user_id": userId,
	})
}
