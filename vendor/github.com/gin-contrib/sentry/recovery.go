package sentry

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
)

// Recovery middleware for sentry crash reporting
func Recovery(client *raven.Client, onlyCrashes bool) gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			flags := map[string]string{
				"endpoint": c.Request.RequestURI,
			}
			if rval := recover(); rval != nil {
				debug.PrintStack()
				rvalStr := fmt.Sprint(rval)
				client.CaptureMessage(rvalStr, flags, raven.NewException(errors.New(rvalStr), raven.NewStacktrace(2, 3, nil)),
					raven.NewHttp(c.Request))
				c.AbortWithStatus(http.StatusInternalServerError)
			}
			if !onlyCrashes {
				for _, item := range c.Errors {
					client.CaptureMessage(item.Error(), flags, &raven.Message{
						Message: item.Error(),
						Params:  []interface{}{item.Meta},
					},
						raven.NewHttp(c.Request))
				}
			}
		}()

		c.Next()
	}
}
