package dbmodels

import (
	"time"

	"gorm.io/datatypes"
)

type ImageMeta struct {
	UUID

	FileID string         `json:"file_id" gorm:"index"`
	Meta   datatypes.JSON `json:"meta"`
	Metas  datatypes.JSON `json:"metas"`

	CreatedAt time.Time `json:"created_at"`
}
