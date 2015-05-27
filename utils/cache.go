package utils

import (
	"bytes"
	"encoding/gob"
	"errors"
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

func SetCache(key string, value interface{}, timeout int64) error {
	data, err := Encode(value)
	if err != nil {
		return err
	}

	err = cc.Put(key, data, timeout)
	if err != nil {
		log.Warnln("Cache失败，key:", key)
		return err
	} else {
		log.Blueln("Cache成功，key:", key)
		return nil
	}
}

func GetCache(key string, to interface{}) error {
	data := cc.Get(key)
	if data == nil {
		return errors.New("Cache不存在")
	}
	// log.Pinkln(data)
	err := Decode(data.([]byte), to)
	if err != nil {
		log.Warnln("获取Cache失败", key, err)
	} else {
		log.Greenln("获取Cache成功", key)
	}

	return err
}

// --------------------
// Encode
// 用gob进行数据编码
//
func Encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// -------------------
// Decode
// 用gob进行数据解码
//
func Decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}
