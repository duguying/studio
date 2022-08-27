package dbmodels

import (
	"duguying/studio/modules/ipip"
	"duguying/studio/service/models"
	"fmt"
	"time"
)

type LoginHistory struct {
	UUID

	UserID    uint       `json:"user_id"`
	SessionID string     `json:"session_id"`
	IP        string     `json:"ip"`
	LoginAt   *time.Time `json:"login_at"`
}

func (lh *LoginHistory) ToModel() *models.LoginHistory {
	area := ""
	loc, err := ipip.GetLocation(lh.IP)
	if err != nil {
		area = "未知"
	} else {
		area = fmt.Sprintf("%s,%s,%s", loc.CountryName, loc.RegionName, loc.CityName)
	}

	return &models.LoginHistory{
		ID:        lh.ID,
		UserID:    lh.UserID,
		SessionID: lh.SessionID,
		IP:        lh.IP,
		Area:      area,
		LoginAt:   lh.LoginAt,
	}
}
