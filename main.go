package main

import (
	"duguying/blog/g"
	"duguying/blog/modules/configuration"
	"duguying/blog/modules/logger"
	"duguying/blog/service"
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	configPath string = "ims.ini"
	logDir     string = "log"
)

func main() {
	versionFlag()

	// 初始化 config
	g.Config = configuration.NewConfig(configPath)

	// 初始化 logger
	initLogger()

	// 初始化 gin
	service.Run()
}

func versionFlag() {
	version := flag.Bool("v", false, "version")
	config := flag.String("c", "ims.ini", "configuration file")
	logDirectory := flag.String("l", "log", "log directory")
	flag.Parse()
	if *version {
		fmt.Println("Git Version: " + g.GitVersion)
		fmt.Println("Build Time: " + g.BuildTime)
		os.Exit(0)
	}

	if *config != "" {
		configPath = *config
	}

	if *logDirectory != "" {
		logDir = *logDirectory
	}
}

func initLogger() {
	expireDefault := time.Hour * 24 * 1
	expireStr := g.Config.Get(logDir, "expire", expireDefault.String())
	expire, err := time.ParseDuration(expireStr)
	if err != nil {
		expire = expireDefault
	}
	logger.InitLogger(logDir, expire)
}
