package dbmodels

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UUID struct {
	Id string `gorm:"type:varchar(40);primary_key;" sql:"comment:'UUID'"`
}

func (b *UUID) BeforeCreate(tx *gorm.DB) error {
	b.Id = uuid.NewV4().String()
	return nil
}
