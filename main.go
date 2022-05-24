package main

import (
	"duguying/studio/docs"
	"duguying/studio/g"
	"duguying/studio/modules/bleve"
	"duguying/studio/modules/configuration"
	"duguying/studio/modules/cron"
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

// @title Studio管理平台API文档
// @version 1.0
// @description This is a Studio Api server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://duguying.net/
// @contact.email rainesli@tencent.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api/v1
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

	// 初始化 swagger
	initSwagger()

	// 初始化 bleve
	bleve.Init()

	// 初始化定时任务
	cron.Init()

	// 初始化 gin
	service.Run(logDir)
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
	path := g.Config.Get("ipip", "path", "/data/ipipfree.ipdb")
	ipip.InitIPIP(path)
}

func initSwagger() {
	listenAddress := g.Config.Get("system", "listen", "127.0.0.1:20192")
	docs.SwaggerInfo.Host = listenAddress
}
