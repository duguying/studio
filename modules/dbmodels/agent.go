package dbmodels

import (
	"duguying/studio/service/models"
	"time"

	"github.com/gogather/json"
)

type Agent struct {
	ID          uint      `json:"id"`
	Online      uint      `json:"online" gorm:"index"` // 1 online, 0 offline
	ClientId    string    `json:"client_id" gorm:"unique;not null"`
	Os          string    `json:"os" gorm:"index"`
	Arch        string    `json:"arch" gorm:"index"`
	Hostname    string    `json:"hostname" gorm:"index"`
	Ip          string    `json:"ip" gorm:"index"`
	IpIns       string    `json:"ip_ins" gorm:"index:,class:FULLTEXT"` // json
	Status      uint      `json:"status" gorm:"index"`
	OnlineTime  time.Time `json:"online_time"`
	OfflineTime time.Time `json:"offline_time"`
}

func (a *Agent) String() string {
	c, _ := json.Marshal(a)
	return string(c)
}

func (a *Agent) ToModel() *models.Agent {
	return &models.Agent{
		ID:          a.ID,
		Online:      a.Online,
		ClientID:    a.ClientId,
		OS:          a.Os,
		Arch:        a.Arch,
		Hostname:    a.Hostname,
		IP:          a.Ip,
		IPIns:       a.IpIns,
		Status:      a.Status,
		OnlineTime:  a.OnlineTime,
		OfflineTime: a.OfflineTime,
	}
}
