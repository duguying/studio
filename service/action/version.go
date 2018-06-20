package action

import (
	"duguying/studio/g"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Version(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"version":     g.Version,
		"git_version": g.GitVersion,
		"build_time":  g.BuildTime,
	})
	return
}

func PageTest(c *gin.Context) {
	fmt.Println("hi")
	c.HTML(http.StatusOK, "test", gin.H{})
}

func PageTest1(c *gin.Context) {
	fmt.Println("hi")
	c.HTML(http.StatusOK, "about/index", gin.H{})
}
