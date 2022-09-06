package models

import "time"

type Agent struct {
	Id          uint
	Online      uint
	ClientId    string
	Os          string
	Arch        string
	Hostname    string
	Ip          string
	Area        string
	IpIns       string
	Status      uint
	OnlineTime  time.Time
	OfflineTime time.Time
}
