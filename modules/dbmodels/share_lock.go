package dbmodels

import "time"

type ShareLock struct {
	LockKey    string    `gorm:"unique;type:varchar(50)"`
	LockOwner  string    `gorm:"index;type:varchar(50)"`
	UpdateTime time.Time `gorm:"index"`
}
