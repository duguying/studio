package action

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Version(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"version": "1.0",
	})
	return
}
