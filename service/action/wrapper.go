// Package action 业务层
package action

import (
	"bytes"
	"duguying/studio/g"
	"duguying/studio/modules/logger"
	"duguying/studio/service/middleware"
	"duguying/studio/service/models"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"duguying/studio/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	"github.com/gogather/json"
)

// CustomContext 自定义web上下文(Context)
type CustomContext struct {
	*gin.Context
}

// HandlerResponseFunc 带响应信息的处理函数
type HandlerResponseFunc func(c *CustomContext) (interface{}, error)

type APILog struct {
	Method string `json:"method"`
	URI    string `json:"uri"`
	Query  string `json:"query"`
	User   string `json:"user"`
	// 如果存在模拟用户的情况,staff记录真实的用户名,user 记录被模拟的用户
	Staff      string `json:"staff"`
	Acquires   string `json:"acquires"`
	Body       string `json:"body" sql:"type:longtext"`
	Response   string `json:"response" sql:"type:longtext"`
	Ok         bool   `json:"ok"`
	Trace      string `json:"trace"`
	ClientIP   string `json:"client_ip"`
	DomainID   string `json:"domain_id,omitempty"`
	ServerID   string `json:"server_id,omitempty"`
	LocationID string `json:"location_id,omitempty"`
	Operator   string `json:"operator"`
	RequestID  string `json:"request_id"`
	Cost       string `json:"cost"`
}

func (al *APILog) ToMap() map[string]interface{} {
	c, _ := json.Marshal(al)
	out := map[string]interface{}{}
	_ = json.Unmarshal(c, &out)
	return out
}

// APIWrapperResponse 带响应信息的api的action包裹器
func APIWrapperResponse(handler HandlerResponseFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				stack := utils.Stack(3)
				stackInfo := fmt.Sprintf("[panic] %v\n%s", err, string(stack))
				c.Set("error_trace", stackInfo)
				fmt.Println("panic error trace:", stackInfo)
				recordPanicReq(c, stackInfo)
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		response, err := handler(&CustomContext{Context: c})
		if err != nil {
			trace := fmt.Sprintf("[err] %s", getErrorTrace(err))
			log.Println("error trace:", trace)
			c.Set("error_trace", trace)
			if response == nil {
				c.JSON(http.StatusOK, models.CommonResponse{
					Ok:  false,
					Msg: err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, response)
			}
			return
		}
		c.JSON(http.StatusOK, response)
		return
	}
}

func recordPanicReq(c *gin.Context, stack string) {
	uri := ""
	u, err := url.ParseRequestURI(c.Request.RequestURI)
	if err != nil {
		uri = c.Request.RequestURI
	} else {
		uri = u.Path
	}

	rl := middleware.RequestLog{
		Method:   c.Request.Method,
		URI:      uri,
		Query:    c.Request.URL.RawQuery,
		Headers:  c.Request.Header,
		ClientIP: c.ClientIP(),
	}

	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logger.L("request").Println("read body error:", err.Error())
	}
	rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
	c.Request.Body = rdr2
	body := string(buf)
	rl.Body = body

	// store api log
	apiLog := &APILog{
		Method:    rl.Method,
		URI:       rl.URI,
		Query:     rl.Query,
		User:      c.GetString("user"),
		Acquires:  c.GetString("acquires"),
		Body:      rl.Body,
		Ok:        false,
		Trace:     stack,
		Operator:  c.GetHeader("X-From"),
		RequestID: c.GetHeader("X-RequestId"),
		ClientIP:  c.ClientIP(),
	}

	g.LogEntry.WithFields(apiLog.ToMap()).Println()
}

func getErrorTrace(err error) (trace string) {
	e, ok := err.(*errors.Error)
	if ok {
		trace = trace + e.ErrorStack()
	} else {
		trace = trace + fmt.Sprintf("%v", err)
	}
	return trace
}
