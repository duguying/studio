package action

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func Version(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"version": "1.0",
	})
	return
}

func PageTest(c *gin.Context) {
	fmt.Println("hi")
	c.HTML(http.StatusOK, "test.tmpl", gin.H{})
}