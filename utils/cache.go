package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/memcache"
	_ "github.com/astaxie/beego/cache/redis"
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
	err := cc.Put(key, value, timeout)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func GetCache(key string) interface{} {
	content := cc.Get(key)
	if content != nil {
		beego.Info("Cache命中, key: " + key)
	}
	return cc.Get(key)
}
