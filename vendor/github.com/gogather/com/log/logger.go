package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Logger struct {
	log.Logger
}

func New(out io.Writer, prefix string, flag int) *Logger {
	return New(out, prefix, flag)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.Logger.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Logger.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (l *Logger) Fatalln(v ...interface{}) {
	l.Logger.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}

func (l *Logger) Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	l.Logger.Output(2, s)
	panic(s)
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.Logger.Output(2, s)
	panic(s)
}

func (l *Logger) Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	l.Logger.Output(2, s)
	panic(s)
}

func (l *Logger) Flags() int {
	return l.Logger.Flags()
}

func (l *Logger) SetFlags(flag int) {
	l.Logger.SetFlags(flag)
}

func (l *Logger) Prefix() string {
	return l.Logger.Prefix()
}

func (l *Logger) SetPrefix(prefix string) {
	l.Logger.SetPrefix(prefix)
}
