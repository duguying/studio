package cache

import (
	ujson "encoding/json"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
	"github.com/gogather/json"
)

// BoltCache 基于bolt实现的kv缓存
type BoltCache struct {
	db         *bolt.DB
	bucketName string
	disabled   bool
}

type boltCacheItem struct {
	Value     string `json:"value"`
	CreatedAt int64  `json:"created_at"`
}

func NewBoltCache() *BoltCache {
	db, err := bolt.Open("cache.db", 0600, &bolt.Options{
		Timeout:         1 * time.Second,
		InitialMmapSize: 1024 * 1024 * 1024 * 1, // 1G
	})
	if err != nil {
		panic(err)
	}
	return NewBolt(db, "session")
}

// NewBolt 新建Bolt缓存实例
func NewBolt(instance *bolt.DB, bucket string) *BoltCache {
	bc := &BoltCache{db: instance, bucketName: bucket}
	err := instance.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bc.bucketName))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return bc
}

// Disable 是否禁用
func (bc *BoltCache) Disable(disable bool) {
	bc.disabled = disable
}

// SetTTL 设置
func (bc *BoltCache) SetTTL(key, value string, expiration time.Duration) error {
	if bc.disabled {
		return nil
	}

	err := bc.db.Update(func(tx *bolt.Tx) error {
		var exp *time.Time = nil
		if expiration > 0 {
			expPoint := time.Now().Add(expiration)
			exp = &expPoint
		}
		item := boltCacheItem{Value: value, CreatedAt: exp.Unix()}
		val, _ := json.Marshal(item)
		b, err := tx.CreateBucketIfNotExists([]byte(bc.bucketName))
		if err != nil {
			return err
		}
		return b.Put([]byte(key), val)
	})
	return err
}

// Get 获取
func (bc *BoltCache) Get(key string) (val string, err error) {
	if bc.disabled {
		return "", fmt.Errorf("bolt cache disabled")
	}

	exist := false
	err = bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bc.bucketName))
		value := b.Get([]byte(key))
		if value == nil {
			exist = false
			val = ""
			return nil
		}

		// check expiration
		item := boltCacheItem{}
		err := ujson.Unmarshal(value, &item)
		if err != nil {
			exist = false
			val = ""
			return nil
		}

		// no expiration
		if item.CreatedAt <= 0 {
			exist = true
			val = item.Value
			return nil
		}

		// get value with calc expiration
		if time.Unix(item.CreatedAt, 0).After(time.Now()) {
			exist = true
			val = item.Value
			return nil
		}

		// expired
		exist = false
		val = ""

		return nil
	})
	if err != nil {
		return "", err
	}
	if exist {
		return val, nil
	} else {
		return val, fmt.Errorf("not exist")
	}
}

// Delete 删除
func (bc *BoltCache) Delete(key string) error {
	if bc.disabled {
		return fmt.Errorf("bolt cache disabled")
	}

	err := bc.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bc.bucketName))
		if err != nil {
			return err
		}
		return b.Delete([]byte(key))
	})
	if err != nil {
		return err
	}
	return nil
}
