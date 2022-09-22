package rlog

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

type lineNumberHook struct {
	callerShortPath bool
}

// Levels 返回所有级别
func (lnh *lineNumberHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire 触发
func (lnh *lineNumberHook) Fire(entry *logrus.Entry) error {
	_, file, line, ok := runtime.Caller(8)
	if ok {
		cl := file
		if lnh.callerShortPath {
			cl = lnh.shortenPath(file)
		}
		entry.Data["file"] = fmt.Sprintf("%s:%d", cl, line)
	}
	return nil
}

// SetShortPath 设置短路径
func (lnh *lineNumberHook) SetShortPath(short bool) {
	lnh.callerShortPath = short
}

func (lnh *lineNumberHook) shortenPath(path string) (shortPath string) {
	hasRoot := strings.HasPrefix(path, "/")
	path = strings.TrimPrefix(path, "/")
	sep := fmt.Sprintf("%c", filepath.Separator)
	segs := strings.Split(path, sep)
	if len(segs) <= 1 {
		return path
	}
	length := len(segs)
	shortSegs := make([]string, length)
	last := segs[length-1]
	for i, seg := range segs {
		if i == length-1 {
			continue
		}
		shortSegs[i] = fmt.Sprintf("%c", seg[0])
	}
	shortSegs[length-1] = last
	shortPath = strings.Join(shortSegs, sep)
	if hasRoot {
		shortPath = sep + shortPath
	}
	return shortPath
}
