package cache

import (
	"log"
	"time"

	"gopkg.in/redis.v5"
)

type RedisCache struct {
	cli *redis.Client
}

func NewRedisCache(cacheOption *CacheRedisOption) *RedisCache {

	redisCli := redis.NewClient(&redis.Options{
		Addr:        cacheOption.Addr,
		Password:    cacheOption.Password,
		DB:          cacheOption.DB,
		PoolSize:    cacheOption.PoolSize,
		ReadTimeout: time.Duration(time.Second * time.Duration(cacheOption.Timeout)),
	})
	err := redisCli.Ping().Err()

	if err != nil {
		log.Println("[system]", err.Error())
	} else {
		log.Println("[system]", "redis connect success")
	}
	return &RedisCache{
		cli: redisCli,
	}
}

const PREFIX = "blog:"
const SESS = "session:"

func (rc *RedisCache) Set(key, value string) error {
	return rc.SetTTL(key, value, 0)
}

func (rc *RedisCache) SetTTL(key, value string, ttl time.Duration) error {
	return rc.cli.Set(PREFIX+key, value, ttl).Err()
}

func (rc *RedisCache) Get(key string) (string, error) {
	return rc.cli.Get(PREFIX + key).Result()
}

func (rc *RedisCache) Delete(key string) error {
	return rc.cli.Del(PREFIX + key).Err()
}

func (rc *RedisCache) SetMapField(key, field string, value interface{}) error {
	return rc.cli.HSet(PREFIX+key, field, value).Err()
}

func (rc *RedisCache) DelMapField(key, field string) error {
	return rc.cli.HDel(PREFIX+key, field).Err()
}

func (rc *RedisCache) GetMap(key string) (map[string]string, error) {
	return rc.cli.HGetAll(PREFIX + key).Result()
}
