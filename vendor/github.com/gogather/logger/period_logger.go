package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

// PeriodLogger a period logger could split log file by day
type PeriodLogger struct {
	gFile       *os.File
	nextLogTime time.Time
	Ls          *log.Logger
	tag         string
	isDefault   bool
	flag        int
}

// NewPeriodLogger new a period logger
func NewPeriodLogger(appName, tag, dir string, isDefault bool, flag int) *PeriodLogger {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		log.Println("mkdir", dir, "err:", err)
		panic(err)
	}

	l := &PeriodLogger{
		nextLogTime: getTonight(),
		tag:         tag,
		isDefault:   isDefault,
		flag:        flag,
	}

	l.setLog(dir, appName)
	go l.logPeriod(dir, appName)

	return l
}

func (l *PeriodLogger) setLog(logDir string, appName string) error {
	yy, mm, dd := l.nextLogTime.Date()

	group := ""
	if !l.isDefault {
		group = fmt.Sprintf("_%s", l.tag)
	}

	filename := fmt.Sprintf("%d_%02d_%02d_%s%s.log", yy, mm, dd, appName, group)
	f, err := os.OpenFile(logDir+"/"+filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	l.gFile = f

	if l.isDefault {
		log.SetOutput(l.gFile)
		log.SetFlags(l.flag)
	}

	l.Ls = log.New(l.gFile, "", l.flag)

	return nil
}

func (l *PeriodLogger) logPeriod(logDir, appName string) {
	for {
		now := time.Now()
		period, _ := days(1)
		if now.After(l.nextLogTime.Add(period)) {
			l.nextLogTime = l.nextLogTime.Add(period)
			oldFile := l.gFile
			err := l.setLog(logDir, appName)
			if err != nil {
				log.Println("set log failed", err, logDir)
			}
			oldFile.Close()
		}
		time.Sleep(time.Second)
	}
}

func getTonight() time.Time {
	now := time.Now()
	date := fmt.Sprintf("%04d-%02d-%02d", now.Year(), now.Month(), now.Day())
	loc, _ := time.LoadLocation("Local")
	pt, _ := time.ParseInLocation("2006-01-02", date, loc)
	return pt
}

func days(d int) (time.Duration, error) {
	return time.ParseDuration(fmt.Sprintf("%dh", d*24))
}
