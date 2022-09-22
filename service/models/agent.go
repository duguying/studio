package models

import "time"

type Agent struct {
	ID          uint      `json:"id"`
	Online      uint      `json:"online"`
	ClientID    string    `json:"client_id"`
	OS          string    `json:"os"`
	Arch        string    `json:"arch"`
	Hostname    string    `json:"hostname"`
	IP          string    `json:"ip"`
	Area        string    `json:"area"`
	IPIns       string    `json:"ip_ins"`
	Status      uint      `json:"status"`
	OnlineTime  time.Time `json:"online_time"`
	OfflineTime time.Time `json:"offline_time"`
}
