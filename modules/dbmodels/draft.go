package dbmodels

import "time"

type Draft struct {
	UUID

	ArticleID uint
	Content   string `gorm:"longtext"`

	CreatedAt time.Time
}
