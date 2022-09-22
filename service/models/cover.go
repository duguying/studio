package models

import "time"

type Cover struct {
	ID            string    `json:"id"`
	FileID        string    `json:"file_id"`
	DayOfWeek     string    `json:"day_of_week"`
	SchedulerType int       `json:"scheduler_type"`
	URL           string    `json:"url"`
	Title         string    `json:"title"`
	CreatedAt     time.Time `json:"created_at"`
}
