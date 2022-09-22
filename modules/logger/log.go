package logger

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/gogather/logger"
)

var gl *logger.GroupLogger

func InitLogger(dir string, expire time.Duration, level int) {
	logSlice := []string{}
	gl = logger.NewGroupLogger(dir, "studio", expire, logSlice, log.Ldate|log.Ltime|log.Lshortfile, level)
}

func L(group string) *logger.Logger {
	return gl.L(group)
}

func GinLogger(logPath string) (io.Writer, error) {
	return os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
}
