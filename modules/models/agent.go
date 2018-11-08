package models

import (
	"github.com/gogather/json"
	"time"
)

type Agent struct {
	Id          uint      `json:"id"`
	Online      uint      `json:"online"` // 1 online, 0 offline
	ClientId    string    `json:"client_id" gorm:"unique;not null"`
	Os          string    `json:"os"`
	Arch        string    `json:"arch"`
	Hostname    string    `json:"hostname"`
	Ip          string    `json:"ip"`
	IpIns       string    `json:"ip_ins"` // json
	Status      uint      `json:"status"`
	OnlineTime  time.Time `json:"online_time"`
	OfflineTime time.Time `json:"offline_time"`
}

func (a *Agent) String() string {
	c, _ := json.Marshal(a)
	return string(c)
}
