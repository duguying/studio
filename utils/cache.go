package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/memcache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/gogather/com"
	"github.com/gogather/com/log"
)

var cc cache.Cache

func InitCache() {
	cacheConfig := beego.AppConfig.String("cache")

	if "redis" == cacheConfig {
		initRedis()
	} else {
		initMemcache()
	}

	log.Greenln("[cache] use", cacheConfig)
}

func initMemcache() {
	var err error
	cc, err = cache.NewCache("memcache", `{"conn":"`+beego.AppConfig.String("memcache_host")+`"}`)

	if err != nil {
		beego.Info(err)
	}

}

func initRedis() {
	var err error
	cc, err = cache.NewCache("redis", `{"conn":"`+beego.AppConfig.String("redis_host")+`"}`)

	if err != nil {
		beego.Info(err)
	}
}

func SetCache(key string, value string, timeout int64) error {
	key = com.Md5(key)
	err := cc.Put(key, value, timeout)
	if err != nil {
		log.Warnln("Cache失败，key:", key)
		return err
	} else {
		log.Blueln("Cache成功，key:", key)
		return nil
	}
}

func GetCache(key string) interface{} {
	key = com.Md5(key)
	content := cc.Get(key)
	if content != nil {
		log.Greenln("Cache命中, key:", key)
	}
	return content
}
