// Package lock 锁
package lock

import "time"

// Lock 锁
type Lock interface {
	GetLock() bool
	ReleaseLock() bool
	LockKey() string
	GetIsLockExtend() bool
	SetIsLockExtend(isLockExtend bool)
	SetLockExtendIntervalSecond(lockExtendIntervalSecond time.Duration)
}
