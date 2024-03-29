// Package cache 缓存
package cache

import (
	"time"
)

type CacheRedisOption struct {
	Timeout  int
	DB       int
	Addr     string
	Password string
	PoolSize int
}

type CacheOption struct {
	Type     string
	Redis    *CacheRedisOption
	BoltPath string
}

// Cache 缓存接口
type Cache interface {
	SetTTL(key string, value string, ttl time.Duration) error
	Set(key string, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

// Init 初始化
func Init(option *CacheOption) Cache {
	var cacheCli Cache
	if option.Type == "redis" {
		cacheCli = NewRedisCache(option.Redis)
	} else {
		cacheCli = NewBoltCache(option.BoltPath)
	}
	return cacheCli
}
