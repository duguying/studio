// Package dbmodels 包注释
package dbmodels

import "time"

type Calendar struct {
	UUID

	Date        time.Time     `json:"date" gorm:"index"`
	Period      time.Duration `json:"period" gorm:"index"`
	Summary     string        `json:"summary" gorm:"type:text"`
	Address     string        `json:"address" gorm:"type:text"`
	Description string        `json:"description" gorm:"type:text"`
	Link        string        `json:"link"`
	Attendee    string        `json:"attendee"`
}
