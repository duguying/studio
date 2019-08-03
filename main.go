package main

import (
	"duguying/studio/g"
	"duguying/studio/modules/configuration"
	"duguying/studio/modules/ipip"
	"duguying/studio/modules/logger"
	"duguying/studio/modules/orm"
	"duguying/studio/modules/redis"
	"duguying/studio/service"
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	configPath string = "studio.ini"
	logDir     string = "log"
)

func main() {
	versionFlag()

	// 初始化 config
	g.Config = configuration.NewConfig(configPath)

	// 初始化 logger
	initLogger()

	// 初始化 ipip
	initIPIP()

	// 初始化 p2p
	//p2p.Init()

	// 初始化 redis
	redis.InitRedisConn()

	// 初始化 database
	orm.InitDatabase()

	// 初始化 gin
	service.Run()
}

func versionFlag() {
	version := flag.Bool("v", false, "version")
	config := flag.String("c", configPath, "configuration file")
	logDirectory := flag.String("l", logDir, "log directory")
	flag.Parse()
	if *version {
		fmt.Println("Version: " + g.Version)
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
	expireStr := g.Config.Get("log", "expire", expireDefault.String())
	expire, err := time.ParseDuration(expireStr)
	if err != nil {
		expire = expireDefault
	}
	level := g.Config.GetInt64("log", "level", 15)
	logger.InitLogger(logDir, expire, int(level))
}

func initIPIP() {
	path := g.Config.Get("ipip", "path", "17monipdb.datx")
	ipip.InitIPIP(path)
}
