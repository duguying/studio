package logger

import (
	"fmt"
	"log"
	"sync"
)

const (
	INFO  = 1 // 0001
	DEBUG = 2 // 0010
	WARN  = 4 // 0100
	ERROR = 8 // 1000
)

// Logger a logger
type Logger struct {
	level int
	d     *log.Logger
	v     *log.Logger
	mu    sync.Mutex
	depth int
}

// NewLogger new a logger
func NewLogger(defaultLogger *log.Logger, v *log.Logger, lv int) *Logger {
	return &Logger{
		d:     defaultLogger,
		v:     v,
		level: lv,
		depth: 2,
	}
}

// set depth
func (l *Logger) SetDepth(depth int) {
	l.depth = depth
}

// Print
func (l *Logger) Println(v ...interface{}) {
	if l.level&INFO <= 0 {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	content := fmt.Sprintln(v...)
	l.d.SetPrefix("[I] ")
	l.d.Output(l.depth, content)
	if l.v != nil {
		l.v.SetPrefix("[I] ")
		l.v.Output(l.depth, content)
	}
}

func (l *Logger) Printf(format string, v ...interface{}) {
	if l.level&INFO <= 0 {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	content := fmt.Sprintf(format, v...)
	l.d.SetPrefix("[I] ")
	l.d.Output(l.depth, content)
	if l.v != nil {
		l.v.SetPrefix("[I] ")
		l.v.Output(l.depth, content)
	}
}

func (l *Logger) Print(v ...interface{}) {
	if l.level&INFO <= 0 {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	content := fmt.Sprint(v...)
	l.d.SetPrefix("[I] ")
	l.d.Output(l.depth, content)
	if l.v != nil {
		l.v.SetPrefix("[I] ")
		l.v.Output(l.depth, content)
	}
}

func (l *Logger) Write(p []byte) (n int, err error) {
	if l.level&INFO <= 0 {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	n, err = l.d.Writer().Write(p)
	if l.v != nil {
		n, err = l.v.Writer().Write(p)
	}
	return n, err
}

// debug

func (l *Logger) Debugln(v ...interface{}) {
	if l.level&DEBUG <= 0 {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	content := fmt.Sprintln(v...)
	l.d.SetPrefix("[D] ")
	l.d.Output(l.depth, content)
	if l.v != nil {
		l.v.SetPrefix("[D] ")
		l.v.Output(l.depth, content)
	}
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.level&DEBUG <= 0 {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	content := fmt.Sprintf(format, v...)
	l.d.SetPrefix("[D] ")
	l.d.Output(l.depth, content)
	if l.v != nil {
		l.v.SetPrefix("[D] ")
		l.v.Output(l.depth, content)
	}
}

func (l *Logger) Debug(v ...interface{}) {
	if l.level&DEBUG <= 0 {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	content := fmt.Sprint(v...)
	l.d.SetPrefix("[D] ")
	l.d.Output(l.depth, content)
	if l.v != nil {
		l.v.SetPrefix("[D] ")
		l.v.Output(l.depth, content)
	}
}

// info

func (l *Logger) Infoln(v ...interface{}) {
	if l.level&INFO <= 0 {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	content := fmt.Sprintln(v...)
	l.d.SetPrefix("[I] ")
	l.d.Output(l.depth, content)
	if l.v != nil {
		l.v.SetPrefix("[I] ")
		l.v.Output(l.depth, content)
	}
}

func (l *Logger) Infof(format string, v ...interface{}) {
	if l.level&INFO <= 0 {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	content := fmt.Sprintf(format, v...)
	l.d.SetPrefix("[I] ")
	l.d.Output(l.depth, content)
	if l.v != nil {
		l.v.SetPrefix("[I] ")
		l.v.Output(l.depth, content)
	}
}

func (l *Logger) Info(v ...interface{}) {
	if l.level&INFO <= 0 {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	content := fmt.Sprint(v...)
	l.d.SetPrefix("[I] ")
	l.d.Output(l.depth, content)
	if l.v != nil {
		l.v.SetPrefix("[I] ")
		l.v.Output(l.depth, content)
	}
}

// warn

func (l *Logger) Warnln(v ...interface{}) {
	if l.level&WARN <= 0 {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	content := fmt.Sprintln(v...)
	l.d.SetPrefix("[W] ")
	l.d.Output(l.depth, content)
	if l.v != nil {
		l.v.SetPrefix("[W] ")
		l.v.Output(l.depth, content)
	}
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.level&WARN <= 0 {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	content := fmt.Sprintf(format, v...)
	l.d.SetPrefix("[W] ")
	l.d.Output(l.depth, content)
	if l.v != nil {
		l.v.SetPrefix("[W] ")
		l.v.Output(l.depth, content)
	}
}

func (l *Logger) Warn(v ...interface{}) {
	if l.level&WARN <= 0 {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	content := fmt.Sprint(v...)
	l.d.SetPrefix("[W] ")
	l.d.Output(l.depth, content)
	if l.v != nil {
		l.v.SetPrefix("[W] ")
		l.v.Output(l.depth, content)
	}
}

// error

func (l *Logger) Errorln(v ...interface{}) {
	if l.level&ERROR <= 0 {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	content := fmt.Sprintln(v...)
	l.d.SetPrefix("[E] ")
	l.d.Output(l.depth, content)
	if l.v != nil {
		l.v.SetPrefix("[E] ")
		l.v.Output(l.depth, content)
	}
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.level&ERROR <= 0 {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	content := fmt.Sprintf(format, v...)
	l.d.SetPrefix("[E] ")
	l.d.Output(l.depth, content)
	if l.v != nil {
		l.v.SetPrefix("[E] ")
		l.v.Output(l.depth, content)
	}
}

func (l *Logger) Error(v ...interface{}) {
	if l.level&ERROR <= 0 {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	content := fmt.Sprint(v...)
	l.d.SetPrefix("[E] ")
	l.d.Output(l.depth, content)
	if l.v != nil {
		l.v.SetPrefix("[E] ")
		l.v.Output(l.depth, content)
	}
}

// fatal

// panic
