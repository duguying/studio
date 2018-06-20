package main

import (
	"duguying/studio/g"
	"duguying/studio/modules/configuration"
	"duguying/studio/modules/ipip"
	"duguying/studio/modules/logger"
	"duguying/studio/modules/orm"
	"duguying/studio/modules/redis"
	"duguying/studio/service"
	"duguying/studio/service/message/store"
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

	// 初始化 redis
	redis.InitRedisConn()

	// 初始化 boltDB
	store.InitBoltDB()

	// 初始化 database
	orm.InitDatabase()

	// 初始化 gin
	service.Run()
}

func versionFlag() {
	version := flag.Bool("v", false, "version")
	config := flag.String("c", "studio.ini", "configuration file")
	logDirectory := flag.String("l", "log", "log directory")
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
	expireStr := g.Config.Get(logDir, "expire", expireDefault.String())
	expire, err := time.ParseDuration(expireStr)
	if err != nil {
		expire = expireDefault
	}
	logger.InitLogger(logDir, expire)
}

func initIPIP() {
	path := g.Config.Get("ipip", "path", "17monipdb.datx")
	ipip.InitIPIP(path)
}
