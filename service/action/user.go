// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/1/25.

package action

import (
	"duguying/studio/g"
	"duguying/studio/modules/db"
	"duguying/studio/modules/session"
	"duguying/studio/service/models"
	"github.com/gin-gonic/gin"
	"github.com/gogather/com"
	"net/http"
	"time"
)

func UserSimpleInfo(c *gin.Context) {

}

// @Router /admin/user_info [get]
// @Tags 用户
// @Description 当前用户信息
// @Success 200 {object} models.UserInfoResponse
func UserInfo(c *gin.Context) {
	userId := uint(c.GetInt64("user_id"))
	user, err := db.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusOK, models.UserInfoResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.UserInfoResponse{
		Ok:   true,
		Data: user.ToInfo(),
	})
	return
}

func UserRegister(c *gin.Context) {
	register := &models.RegisterArgs{}
	err := c.BindJSON(register)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"msg": err.Error(),
		})
		return
	}
	user, err := db.RegisterUser(register.Username, register.Password, register.Email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"msg": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"ok": true,
			"user": gin.H{
				"id":       user.Id,
				"username": user.Username,
			},
		})
		return
	}
}

// @Router /user_login [put]
// @Tags 用户
// @Description 用户登录
// @Param auth body models.LoginArgs true "登录鉴权信息"
// @Success 200 {object} models.LoginResponse
func UserLogin(c *gin.Context) {
	login := &models.LoginArgs{}
	err := c.BindJSON(login)
	if err != nil {
		c.JSON(http.StatusOK, models.LoginResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}
	user, err := db.GetUser(login.Username)
	if err != nil {
		c.JSON(http.StatusOK, models.LoginResponse{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}

	// validate
	passwd := com.Md5(login.Password + user.Salt)
	if passwd != user.Password {
		c.JSON(http.StatusOK, models.LoginResponse{
			Ok:  false,
			Msg: "login failed, invalid password",
		})
		return
	} else {
		sid := session.SessionID()
		if sid == "" {
			c.JSON(http.StatusOK, models.LoginResponse{
				Ok:  false,
				Msg: "generate sid failed",
			})
			return
		} else {
			defaultSessionTime := time.Hour * 24
			sessionTimeCfg := g.Config.Get("session", "expire", defaultSessionTime.String())
			sessionExpire, err := time.ParseDuration(sessionTimeCfg)
			if err != nil {
				c.JSON(http.StatusOK, models.LoginResponse{
					Ok:  false,
					Msg: err.Error(),
				})
				return
			} else {
				// store session
				session.SessionSet(sid, sessionExpire, &session.Entity{
					UserId: user.Id,
				})

				c.JSON(http.StatusOK, models.LoginResponse{
					Ok:  true,
					Sid: sid,
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
		"msg":     "logout success",
		"user_id": userId,
	})
}

func UsernameCheck(c *gin.Context) {
	username := c.DefaultQuery("username", "")
	if username == "" {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"msg": "username could not be empty",
		})
		return
	} else {
		valid, err := db.CheckUsername(username)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"ok":  false,
				"msg": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"ok":    true,
				"valid": valid,
			})
			return
		}
	}
}
