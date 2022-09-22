package dbmodels

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UUID struct {
	ID string `gorm:"type:varchar(40);primary_key;" sql:"comment:'UUID'"`
}

func (b *UUID) BeforeCreate(tx *gorm.DB) error {
	b.ID = uuid.New().String()
	return nil
}
