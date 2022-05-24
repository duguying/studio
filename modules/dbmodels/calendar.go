// Package dbmodels 包注释
package dbmodels

import "time"

type Calendar struct {
	UUID

	Start       time.Time  `json:"start"`
	End         time.Time  `json:"end"`
	Stamp       time.Time  `json:"stamp"`
	Summary     string     `json:"summary"`
	Address     string     `json:"address"`
	Description string     `json:"description"`
	Link        string     `json:"link"`
	Attendee    string     `json:"attendee"`
	SendAt      *time.Time `json:"send_at"`
}
