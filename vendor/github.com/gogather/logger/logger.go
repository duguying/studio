package logger

import (
	"log"
	"fmt"
)

// GROUP option for enable group log
var GROUP = true

// Logger a logger
type Logger struct {
	d *log.Logger
	v *log.Logger
}

// NewLogger new a logger
func NewLogger(defaultLogger *log.Logger, v *log.Logger) *Logger {
	return &Logger{
		d: defaultLogger,
		v: v,
	}
}

// Println print line like log.Println
func (l *Logger) Println(v ...interface{}) {
	l.d.Output(2, fmt.Sprintln(v...))
	if GROUP {
		l.v.Output(2, fmt.Sprintln(v...))
	}
}

// Printf print format like log.Printf
func (l *Logger) Printf(format string, v ...interface{}) {
	l.d.Output(2, fmt.Sprintf(format, v...))
	if GROUP && l.v != nil {
		l.v.Output(2, fmt.Sprintf(format, v...))
	}
}
