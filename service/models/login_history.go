package models

import "time"

type LoginHistory struct {
	ID        string     `json:"id"`
	UserID    uint       `json:"user_id"`
	SessionID string     `json:"session_id"`
	IP        string     `json:"ip"`
	Area      string     `json:"area"`
	Expired   bool       `json:"expired"`
	LoginAt   *time.Time `json:"login_at"`
}
