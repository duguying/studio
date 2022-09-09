package dbmodels

import (
	"time"

	"gorm.io/datatypes"
)

type ImageMeta struct {
	UUID

	FileID string         `json:"file_id" gorm:"index"`
	Meta   datatypes.JSON `json:"meta" gorm:"index:,class:FULLTEXT"`
	Metas  datatypes.JSON `json:"metas" gorm:"index:,class:FULLTEXT"`

	CreatedAt time.Time `json:"created_at"`
}
