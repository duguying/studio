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
	Type  string
	Redis *CacheRedisOption
}

type Cache interface {
	SetTTL(key string, value string, ttl time.Duration) error
	Get(key string) (string, error)
	Delete(key string) error
}

func Init(option *CacheOption) Cache {
	var cacheCli Cache
	if option.Type == "redis" {
		cacheCli = NewRedisCache(option.Redis)
	} else {
		cacheCli = NewBoltCache()
	}
	return cacheCli
}
