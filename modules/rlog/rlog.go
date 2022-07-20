// Package rlog rlog
package rlog

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/dogenzaka/rotator"
	"github.com/gogather/com"
	"github.com/sirupsen/logrus"
)

type RemoteAdaptor interface {
	Report(interface{}) error
	Close()
}

type chanWriter struct {
	logChan    chan string
	allowBlock bool
}

// NewChanWriter 创建chanWriter
func NewChanWriter(logChan chan string, allowBlock bool) *chanWriter {
	return &chanWriter{logChan: logChan, allowBlock: allowBlock}
}

// Write 写
func (c *chanWriter) Write(p []byte) (n int, err error) {
	if c.allowBlock {
		// 阻塞，不建议开启：本地日志有且level小于cfg或者有业务逻辑，消费速度基本足够，有丢失可考虑增加obj数量
		c.logChan <- string(p)
	} else {
		// 非阻塞，可能丢数据，避免消费性能不足影响业务逻辑
		select {
		case c.logChan <- string(p):
		default:
		}
	}

	return len(p), nil
}

// RLog RLog
type RLog struct {
	remoteAddr       string
	remoteCli        RemoteAdaptor
	remoteSendThread int
	logrusInstance   *logrus.Logger
	logChan          chan string
	writer           *chanWriter
	enableRemote     bool
	logFilePath      string
	rotatorFile      *rotator.SizeRotator
	myIP             string
}

// NewRLog 创建RLog
func NewRLog(ctx context.Context, topic string, path string) *RLog {
	rl := &RLog{
		enableRemote:     true,
		remoteSendThread: 4,
		logFilePath:      path,
	}
	rl.initRotatorLog()

	rl.logrusInstance = logrus.New()
	rl.logrusInstance.WithContext(ctx)
	rl.logChan = make(chan string, 100000)
	rl.writer = NewChanWriter(rl.logChan, false)

	multiWriter := io.MultiWriter(rl.writer, rl.rotatorFile)
	rl.logrusInstance.SetOutput(multiWriter)
	rl.logrusInstance.SetFormatter(&logrus.JSONFormatter{})
	rl.logrusInstance.SetReportCaller(false)

	lineNumHook := &lineNumberHook{}
	lineNumHook.SetShortPath(true)
	rl.logrusInstance.AddHook(lineNumHook)

	if rl.enableRemote {
		rl.initRemoteClient(topic)
	}
	rl.initChanReader()

	return rl
}

// WithFields 添加属性
func (rl *RLog) WithFields(fields map[string]interface{}) (entry *logrus.Entry) {
	return rl.logrusInstance.WithFields(fields)
}

func (rl *RLog) initRotatorLog() {
	err := com.WriteFileAppendWithCreatePath(rl.logFilePath, "")
	if err != nil {
		fmt.Println("create local log file failed, err:", err)
		return
	}
	rl.rotatorFile = rotator.NewSizeRotator(rl.logFilePath)
	rl.rotatorFile.MaxRotation = 99                      // 99 files
	rl.rotatorFile.RotationSize = int64(1024 * 1024 * 1) // 100M
}

// Close 关闭日志
func (rl *RLog) Close() {
	close(rl.logChan)
	rl.closeRotatorLog()
	if rl.enableRemote {
		rl.closeRemote()
	}
}

func (rl *RLog) closeRotatorLog() {
	if rl.rotatorFile != nil {
		err := rl.rotatorFile.Close()
		if err != nil {
			fmt.Println("close rotator file failed, err:", err.Error())
		}
	}
}

func (rl *RLog) closeRemote() {
	rl.remoteCli.Close()
}

// SetZhiyanEnable 设置开启Zhiyan日志
func (rl *RLog) SetZhiyanEnable(enable bool) {
	rl.enableRemote = enable
	if !enable {
		rl.remoteCli = nil
	}
}

func (rl *RLog) initChanReader() {
	for i := 0; i < rl.remoteSendThread; i++ {
		go func() {
			for {
				select {
				case line := <-rl.logChan:
					{
						rl.send(line)
					}
				}
			}
		}()
	}
}

func (rl *RLog) send(line string) {
	err := rl.sendRemoteMessage(line)
	if err != nil {
		fmt.Println("send remote failed, err:", err.Error())
	}
}

func (rl *RLog) initRemoteClient(topic string) (err error) {
	myIP := rl.getIPAddr()
	rl.myIP = myIP

	rl.remoteCli, err = NewEsAdaptor(rl.remoteAddr, topic)
	if err != nil {
		return err
	}
	return nil
}

func (rl *RLog) sendRemoteMessage(msg string) error {
	if rl.remoteCli != nil {
		return rl.remoteCli.Report(msg)
	}
	return nil
}

func (rl *RLog) getIPAddr() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
