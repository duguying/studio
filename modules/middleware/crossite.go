package middleware

import (
	"duguying/studio/g"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServerMark() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Server", fmt.Sprintf("duguying.net - %s", g.GitVersion))
		c.Next()
	}
}

func CrossSite() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		c.Writer.Header().Set("Vary", "Origin")
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT, GET, OPTIONS, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, X-CSRF-TOKEN, X-Token")

		if c.Request.Method == "OPTIONS" {
			c.Status(http.StatusOK)
			c.Abort()
		} else {
			c.Next()
		}
	}
}
