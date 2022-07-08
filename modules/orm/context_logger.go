package orm

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

// Colors
const (
	Reset       = "\033[0m"
	Red         = "\033[31m"
	Green       = "\033[32m"
	Yellow      = "\033[33m"
	Blue        = "\033[34m"
	Magenta     = "\033[35m"
	Cyan        = "\033[36m"
	White       = "\033[37m"
	MagentaBold = "\033[35;1m"
	RedBold     = "\033[31;1m"
	YellowBold  = "\033[33;1m"
)

type Config struct {
	SlowThreshold time.Duration
	Colorful      bool
	LogLevel      logger.LogLevel
}

var Default = New(Config{
	SlowThreshold: 100 * time.Millisecond,
	LogLevel:      logger.Warn,
	Colorful:      true,
})

type slogger struct {
	Config
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

// LogMode log mode
func (l *slogger) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

func (l slogger) Printf(ctx context.Context, format string, args ...interface{}) {
	reqPrefix := ""
	reqid, ok := ctx.Value("reqid").(string)
	if ok {
		reqPrefix = fmt.Sprintf("[%s] ", reqid)
	}

	logseg := fmt.Sprintf(format, args...)
	buf, ok := ctx.Value("lbw").(*bytes.Buffer)
	if ok {
		buf.WriteString(logseg + "\n")
	}

	sqlog := ""
	segs := strings.Split(logseg, "\n")
	for _, seg := range segs {
		sqlog = sqlog + fmt.Sprintln(reqPrefix+seg)
	}

	fmt.Print(sqlog)
}

// Info print info
func (l slogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		l.Printf(ctx, l.infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Warn print warn messages
func (l slogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		l.Printf(ctx, l.warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Error print error messages
func (l slogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		l.Printf(ctx, l.errStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Trace print sql message
func (l slogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel > 0 {
		elapsed := time.Since(begin)
		switch {
		case err != nil && l.LogLevel >= logger.Error:
			sql, rows := fc()
			l.Printf(ctx, l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
			sql, rows := fc()
			l.Printf(ctx, l.traceWarnStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		case l.LogLevel >= logger.Info:
			sql, rows := fc()
			l.Printf(ctx, l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}

func New(config Config) logger.Interface {
	var (
		infoStr      = "%s [info] "
		warnStr      = "%s [warn] "
		errStr       = "%s [error] "
		traceStr     = "%s [%v] [rows:%d] %s"
		traceWarnStr = "%s [%v] [rows:%d] %s"
		traceErrStr  = "%s %s [%v] [rows:%d] %s"
	)

	if config.Colorful {
		infoStr = Green + "%s\n" + Reset + Green + "[info] " + Reset
		warnStr = Blue + "%s\n" + Reset + Magenta + "[warn] " + Reset
		errStr = Magenta + "%s\n" + Reset + Red + "[error] " + Reset
		traceStr = Green + "%s\n" + Reset + Yellow + "[%.3fms] " + Blue + "[rows:%d]" + Reset + " %s"
		traceWarnStr = Green + "%s\n" + Reset + RedBold + "[%.3fms] " + Yellow + "[rows:%d]" + Magenta + " %s" + Reset
		traceErrStr = RedBold + "%s " + MagentaBold + "%s\n" + Reset + Yellow + "[%.3fms] " + Blue + "[rows:%d]" + Reset + " %s"
	}

	return &slogger{
		// Writer:       writer,
		Config:       config,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
}
