package dbmodels

import (
	"duguying/studio/service/models"
	"time"
)

const (
	SchedulerTypeDefault = 0
	SchedulerTypeDay     = 1
)

type Cover struct {
	UUID

	FileID        string    `json:"file_id"`
	DayOfWeek     string    `json:"day_of_week"`
	SchedulerType int       `json:"scheduler_type"`
	Title         string    `json:"title"`
	CreatedAt     time.Time `json:"created_at"`
}

func (c *Cover) ToModel() *models.Cover {
	return &models.Cover{
		ID:            c.ID,
		FileID:        c.FileID,
		DayOfWeek:     c.DayOfWeek,
		SchedulerType: c.SchedulerType,
		Title:         c.Title,
		CreatedAt:     c.CreatedAt,
	}
}
