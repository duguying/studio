package dbmodels

import "time"

type Draft struct {
	UUID

	ArticleID uint   `gorm:"index"`
	Content   string `gorm:"longtext"`

	CreatedAt time.Time
}
