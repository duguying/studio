package logger

import (
	"github.com/gogather/cleaner"
	"io"
	"os"
	"path/filepath"
	"time"
)

// GroupLogger print group period log
type GroupLogger struct {
	loggerMap     map[string]*PeriodLogger
	defaultLogger *PeriodLogger
	expire        time.Duration
	flag          int
	level         int
	excludes      []string
}

// NewGroupLogger new a group logger manager
func NewGroupLogger(dir string, appName string, expire time.Duration, logSlice []string, flag int, lv int) *GroupLogger {
	gl := &GroupLogger{
		loggerMap:     map[string]*PeriodLogger{},
		defaultLogger: NewPeriodLogger(appName, "", dir, true, flag, lv),
		expire:        expire,
		flag:          flag,
		level:         lv,
	}

	for i := 0; i < len(logSlice); i++ {
		sliceName := logSlice[i]
		if len(sliceName) <= 0 {
			continue
		}
		downloadLogger := NewPeriodLogger(appName, sliceName, filepath.Join(dir, sliceName), false, flag, lv)
		gl.loggerMap[sliceName] = downloadLogger
	}

	if expire < time.Hour*24 {
		gl.expire = time.Hour * 24
	}

	fc := cleaner.New(dir, gl.expire, time.Hour, -1)
	fc.SetFilter(func(path string, info os.FileInfo) (willClean bool) {
		for _, excludePath := range gl.excludes {
			if excludePath == path {
				return false
			}
		}
		return true
	})
	fc.StartCleanTask()

	return gl
}

// set excludes path
func (gl *GroupLogger) SetExcludes(excludePath ...string) {
	gl.excludes = excludePath
}

// default logger
func (gl *GroupLogger) D() *Logger {
	d := gl.defaultLogger
	return NewLogger(d.Ls, nil, gl.level)
}

// L get the logger from group logger
func (gl *GroupLogger) L(group string) *Logger {
	d := gl.defaultLogger
	v, ok := gl.loggerMap[group]
	if ok {
		return NewLogger(d.Ls, v.Ls, gl.level)
	}
	return NewLogger(d.Ls, nil, gl.level)
}

// W get the log writer
func (gl *GroupLogger) W(group string) io.Writer {
	d := gl.defaultLogger
	v, ok := gl.loggerMap[group]
	if ok {
		return v.gFile
	}
	return d.gFile
}
