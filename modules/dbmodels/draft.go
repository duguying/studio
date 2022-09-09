package dbmodels

import "time"

type Draft struct {
	UUID

	ArticleID uint   `gorm:"index"`
	Content   string `gorm:"type:longtext;index:,class:FULLTEXT"`

	CreatedAt time.Time
}
