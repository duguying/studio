// Copyright 2019. All rights reserved.
// This file is part of sparta-admin project
// I am coding in Tencent
// Created by rainesli on 2019/3/19.

// Package middleware gin中间件
package middleware

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"duguying/studio/g"
	"duguying/studio/service/models"

	"github.com/gin-gonic/gin"
	"github.com/gogather/json"
	"github.com/google/uuid"
)

// RequestLog 请求日志结构
type RequestLog struct {
	Method   string      `json:"method"`
	URI      string      `json:"uri"`
	Query    string      `json:"query"`
	Headers  http.Header `json:"headers"`
	Body     string      `json:"body"`
	ClientIP string      `json:"client_ip"`
}

// String 序列化
func (rl *RequestLog) String() string {
	c, _ := json.Marshal(rl)
	return string(c)
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write 写入
func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// ResponseLog 响应日志
type ResponseLog struct {
	Status int    `json:"status"`
	Ok     bool   `json:"ok"`
	Msg    string `json:"msg"`
}

// String 序列化
func (rlog *ResponseLog) String() string {
	c, _ := json.Marshal(rlog)
	return string(c)
}

// RestLog Restful接口日志中间件
func RestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		startTimeMs := startTime.UnixNano() / int64(time.Millisecond)
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		uri := ""
		u, err := url.ParseRequestURI(c.Request.RequestURI)
		if err != nil {
			uri = c.Request.RequestURI
		} else {
			uri = u.Path
		}
		if skipURI(uri) {
			c.Next()
			return
		}
		reqID := uuid.New().String()
		if c.GetHeader("X-RequestId") != "" {
			reqID = c.GetHeader("X-RequestId")
		}
		c.Header("X-RequestId", reqID)
		c.Set("X-RequestId", reqID)
		rl := RequestLog{
			Method:   c.Request.Method,
			URI:      uri,
			Query:    c.Request.URL.RawQuery,
			Headers:  c.Request.Header,
			ClientIP: c.ClientIP(),
		}
		buf, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			g.LogEntry.WithField("slice", "request").Println("read body error:", err.Error())
		}
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
		c.Request.Body = rdr2
		body := string(buf)
		rl.Body = body
		g.LogEntry.WithField("slice", "request").Println("request:", rl.String())
		c.Next()
		statusCode := c.Writer.Status()
		if isMethodRecord(rl.Method) {
			rlog := &ResponseLog{}
			rawBytes := blw.body.Bytes()
			rsp := string(rawBytes)
			err = json.Unmarshal(rawBytes, rlog)
			if err != nil {
				g.LogEntry.WithField("slice", "request").Println("parse response failed, err:", err.Error(), "raw:", string(rawBytes))
				if len(rawBytes) > 1024*512 {
					rsp = fmt.Sprintf("[len:%d]", len(rawBytes))
				}
			} else {
				rlog.Status = statusCode
				g.LogEntry.WithField("slice", "request").Println("response:", rlog.String())
			}

			apiLog := &models.APILog{
				Method:    rl.Method,
				URI:       rl.URI,
				Query:     rl.Query,
				Body:      rl.Body,
				Response:  rsp,
				Ok:        rlog.Ok,
				RequestID: reqID,
				ClientIP:  c.ClientIP(),
				CreatedAt: time.Now(),
				Cost:      time.Since(startTime).String(),
			}
			if g.Db.Dialector.Name() == "sqlite3" {
			} else {
				go recordLog(c, apiLog, startTimeMs)
			}
		}
	}
}

func recordLog(c *gin.Context, logInfo *models.APILog, startTime int64) {
	finishTime := time.Now().UnixNano() / int64(time.Millisecond)
	cost := finishTime - startTime
	g.LogEntry.WithFields(logInfo.ToMap()).
		Printf("cost: %s", (time.Duration(cost) * time.Millisecond).String())
}

func isMethodRecord(method string) bool {
	if g.Config.Get("api-log", "write-only", "false") == "true" {
		if method != http.MethodGet && method != http.MethodOptions {
			return true
		} else {
			return false
		}
	} else {
		return true
	}
}

func skipURI(uri string) bool {
	skipURIMap := g.Config.GetSectionAsMap("skip-uri-log")
	val, ok := skipURIMap[uri]
	if ok && val == "enable" {
		return true
	}
	return false
}
