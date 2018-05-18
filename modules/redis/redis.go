package redis

import (
	"gopkg.in/redis.v5"
	"log"
	"duguying/blog/g"
	"strconv"
	"time"
)

func InitRedisConn() {
	readTimeout, _ := strconv.Atoi(g.Config.Get("redis", "timeout", "4"))
	db, _ := strconv.Atoi(g.Config.Get("redis", "db", "11"))
	g.Redis = redis.NewClient(&redis.Options{
		Addr:        g.Config.Get("redis", "addr", ""),
		Password:    g.Config.Get("redis", "password", ""),
		DB:          db,
		PoolSize:    int(g.Config.GetInt64("redis","pool-size",1000)),
		ReadTimeout: time.Duration(time.Second * time.Duration(readTimeout)),
	})
	err := g.Redis.Ping().Err()

	if err != nil {
		log.Println("[system]", err.Error())
	} else {
		log.Println("[system]", "redis connect success")
	}
}

const PREFIX = "blog:"

func Set(key, value string) error {
	return SetTTL(key, value, 0)
}

func SetTTL(key, value string, ttl time.Duration) error {
	return g.Redis.Set(PREFIX+key, value, ttl).Err()
}

func Get(key string) (string, error) {
	return g.Redis.Get(PREFIX + key).Result()
}

func Delete(key string) error {
	return g.Redis.Del(PREFIX + key).Err()
}

func SetMapField(key, field string, value interface{}) error {
	return g.Redis.HSet(PREFIX+key, field, value).Err()
}

func DelMapField(key, field string) error {
	return g.Redis.HDel(PREFIX+key, field).Err()
}

func GetMap(key string) (map[string]string, error) {
	return g.Redis.HGetAll(PREFIX + key).Result()
}
