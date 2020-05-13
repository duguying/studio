package dbmodels

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type UUID struct {
	Id string `gorm:"type:varchar(40);primary_key;" sql:"comment:'UUID'"`
}

func (b *UUID) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.NewV4()
	return scope.SetColumn("Id", id.String())
}
