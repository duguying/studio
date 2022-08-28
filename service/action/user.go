// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/1/25.

package action

import (
	"duguying/studio/g"
	"duguying/studio/modules/db"
	"duguying/studio/modules/session"
	"duguying/studio/service/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gogather/com"
)

// UserSimpleInfo 用户简单信息
// @Router /admin/user_info [get]
// @Tags 用户
// @Description 当前用户信息
// @Success 200 {object} models.UserInfoResponse
func UserSimpleInfo(c *CustomContext) (interface{}, error) {
	return models.CommonResponse{Ok: true}, nil
}

// UserInfo 用户信息
// @Router /admin/user_info [get]
// @Tags 用户
// @Description 当前用户信息
// @Success 200 {object} models.UserInfoResponse
func UserInfo(c *gin.Context) {
	userID := uint(c.GetInt64("user_id"))
	user, err := db.GetUserById(g.Db, userID)
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
	user, err := db.RegisterUser(g.Db, register.Username, register.Password, register.Email)
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

// UserLogin 用户登录
// @Router /user_login [put]
// @Tags 用户
// @Description 用户登录
// @Param auth body models.LoginArgs true "登录鉴权信息"
// @Success 200 {object} models.LoginResponse
func UserLogin(c *CustomContext) (interface{}, error) {
	login := &models.LoginArgs{}
	err := c.BindJSON(login)
	if err != nil {
		return nil, err
	}
	user, err := db.GetUser(g.Db, login.Username)
	if err != nil {
		return nil, err
	}

	// validate
	passwd := com.Md5(login.Password + user.Salt)
	if passwd != user.Password {
		return nil, fmt.Errorf("登陆失败，密码错误")
	} else {
		sid := session.SessionID()
		if sid == "" {
			return nil, fmt.Errorf("生成会话失败")
		} else {
			defaultSessionTime := time.Hour * 24
			sessionTimeCfg := g.Config.Get("session", "expire", defaultSessionTime.String())
			sessionExpire, err := time.ParseDuration(sessionTimeCfg)
			if err != nil {
				return nil, err
			} else {
				// store session
				entity := &session.Entity{
					UserID:    user.Id,
					IP:        c.ClientIP(),
					LoginAt:   time.Now(),
					UserAgent: c.Request.UserAgent(),
				}
				session.SessionSet(sid, sessionExpire, entity)

				err = db.AddLoginHistory(g.Db, sid, entity)
				if err != nil {
					return nil, err
				}

				return models.LoginResponse{
					Ok:  true,
					Sid: sid,
				}, nil
			}

		}
	}
}

func UserLogout(c *gin.Context) {
	sid := c.GetString("sid")
	userID := uint(c.GetInt64("user_id"))
	session.SessionDel(sid)
	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"msg":     "logout success",
		"user_id": userID,
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
		valid, err := db.CheckUsername(g.Db, username)
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

// ListUserLoginHistory 登陆历史列表
// @Router admin/login_history [get]
// @Tags 用户
// @Description 登陆历史列表
// @Param page query uint true "页码"
// @Param size query uint true "每页数"
// @Success 200 {object} models.ListUserLoginHistoryResponse
func ListUserLoginHistory(c *CustomContext) (interface{}, error) {
	userID := uint(c.UserID())
	req := models.CommonPagerRequest{}
	err := c.BindQuery(&req)
	if err != nil {
		return nil, err
	}

	list, total, err := db.PageLoginHistoryByUserID(g.Db, userID, req.Page, req.Size)
	if err != nil {
		return nil, err
	}

	apiList := []*models.LoginHistory{}
	for _, item := range list {
		hist := item.ToModel()
		entity := session.SessionGet(hist.SessionID)
		if entity != nil {
			hist.Expired = false
		} else {
			hist.Expired = true
		}
		apiList = append(apiList, hist)
	}

	return models.ListUserLoginHistoryResponse{
		Ok:    true,
		Total: int(total),
		List:  apiList,
	}, nil
}

func UserMessageCount(c *CustomContext) (interface{}, error) {
	return gin.H{
		"ok": true,
		"data": map[string]interface{}{
			"count": 0,
		},
	}, nil
}
