package lock

import (
	"duguying/studio/modules/dbmodels"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// DBShareLock 数据库共享锁
type DBShareLock struct {
	lockKey                  string
	lockOwner                string
	isLock                   bool
	lockTimeoutSecond        time.Duration
	isLockExtend             bool // 表示是否定期延长锁
	lockExtendIntervalSecond time.Duration
	db                       *gorm.DB
}

// MakeDBLock 创建数据库共享锁
func MakeDBLock(db *gorm.DB, lockKey string, lockTimeoutSecond time.Duration) (shareLock *DBShareLock) {
	var lockOwner = uuid.New().String()
	localIp := getIpAddr()
	if localIp == "" {
		log.Println("localIp is empty!")
	} else {
		lockOwner = localIp + "_" + uuid.New().String()
	}
	shareLock = &DBShareLock{
		lockKey:                  lockKey,
		lockOwner:                lockOwner,
		isLock:                   false,
		lockTimeoutSecond:        lockTimeoutSecond,
		db:                       db,
		isLockExtend:             false,                 // 默认不定期延长锁时间
		lockExtendIntervalSecond: lockTimeoutSecond / 2, // 默认每隔超时间的一半就把锁时间延长
	}
	return shareLock
}

// GetIsLockExtend 获取IsLockExtend
func (shareLock *DBShareLock) GetIsLockExtend() bool {
	return shareLock.isLockExtend
}

// SetIsLockExtend 设置isLockExtend
func (shareLock *DBShareLock) SetIsLockExtend(isLockExtend bool) {
	shareLock.isLockExtend = isLockExtend
}

// SetLockExtendIntervalSecond 设置Interval
func (shareLock *DBShareLock) SetLockExtendIntervalSecond(lockExtendIntervalSecond time.Duration) {
	shareLock.lockExtendIntervalSecond = lockExtendIntervalSecond
}

// GetLock 获取锁
func (shareLock *DBShareLock) GetLock() bool {
	succ := getShareLock(shareLock)
	if succ {
		shareLock.isLock = true
	}
	if shareLock.isLock && shareLock.isLockExtend {
		// 后台更新锁时间
		go func() {
			for {
				time.Sleep(shareLock.lockExtendIntervalSecond)
				if !shareLock.isLock || !shareLock.isLockExtend {
					// log.Printf("stop update lock update_time\n")
					break
				}
				// log.Printf("update lock update_time\n")
				updateShareLockTime(shareLock)
			}
		}()
	}
	return succ
}

// ReleaseLock 释放锁
func (shareLock *DBShareLock) ReleaseLock() bool {
	if releaseShareLock(shareLock) {
		shareLock.isLock = false
		return true
	} else {
		return false
	}
}

// LockKey 返回Lock Key
func (shareLock *DBShareLock) LockKey() string {
	return shareLock.lockKey
}

func getShareLock(shareLock *DBShareLock) bool {
	// 查找尚未过期的锁
	exLock := &dbmodels.ShareLock{}
	timeoutTime := time.Now().Add(-1 * shareLock.lockTimeoutSecond)
	err := shareLock.db.Model(dbmodels.ShareLock{}).
		Where("lock_key=? and update_time>?", shareLock.lockKey, timeoutTime).First(exLock).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			// 没有现成的未过期的锁，尝试插入，注意，时间更新为当前的
			// 先删除过期的
			err := shareLock.db.Model(dbmodels.ShareLock{}).
				Where("lock_key=? and update_time<?", shareLock.lockKey, timeoutTime).
				Delete(&dbmodels.ShareLock{}).Error
			if err != nil {
				log.Printf("delete lock err=%s\n", err.Error())
				return false
			}
			err = shareLock.db.Model(dbmodels.ShareLock{}).Create(&dbmodels.ShareLock{
				LockKey:    shareLock.lockKey,
				LockOwner:  shareLock.lockOwner,
				UpdateTime: time.Now(),
			}).Error
			if err != nil {
				// 因为设置了unique约束，可能获得锁失败
				return false
			} else {
				// 获得锁成功
				return true
			}
		} else {
			log.Println("getShareLock err=", err.Error())
			return false
		}
	} else {
		// 有现成的锁，检查owner是否是它
		if exLock.LockOwner == shareLock.lockOwner {
			return true
		} else {
			return false
		}
	}
}

func releaseShareLock(shareLock *DBShareLock) bool {
	err := shareLock.db.Model(dbmodels.ShareLock{}).
		Where("lock_key=? and lock_owner=?", shareLock.lockKey, shareLock.lockOwner).
		Delete(&dbmodels.ShareLock{}).Error
	if err != nil {
		log.Printf("try delete lock error,%v\n", err)
		return false
	}
	return true
}

func updateShareLockTime(shareLock *DBShareLock) {
	// 定时更新锁
	err := shareLock.db.Model(dbmodels.ShareLock{}).
		Where("lock_key=? and lock_owner=?", shareLock.lockKey, shareLock.lockOwner).
		Updates(map[string]interface{}{
			"update_time": time.Now(),
		}).Error
	if err != nil {
		log.Printf("update lock err=%s\n", err.Error())
	}
}

func getIpAddr() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}

		}
	}
	return ""
}
