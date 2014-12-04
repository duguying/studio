package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/memcache"
	// "log"
)

var cc cache.Cache

func InitCache() {
	var err error
	cc, err = cache.NewCache("memcache", `{"conn":"`+beego.AppConfig.String("memcache_host")+`"}`)

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
