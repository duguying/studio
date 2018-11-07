package logger

import (
	"github.com/gogather/logger"
	"io"
	"log"
	"os"
	"time"
)

var gl *logger.GroupLogger

func InitLogger(dir string, expire time.Duration) {
	logSlice := []string{"blog", "orm", "ws", "browserrcv", "browsersnt"}
	gl = logger.NewGroupLogger(dir, "studio", expire, logSlice, log.Ldate|log.Ltime|log.Lshortfile)
}

func L(group string) *logger.Logger {
	return gl.L(group)
}

func GinLogger(logPath string) (io.Writer, error) {
	return os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
}
