package dbmodels

import "duguying/studio/service/models"

type TrojanUsers struct {
	ID       uint   `json:"id"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Quota    int64  `json:"quota"`
	Download int64  `json:"download"`
	Upload   int64  `json:"upload"`
}

func (tu *TrojanUsers) TableName() string {
	return "users"
}

func (tu *TrojanUsers) ToModel() *models.TrojanUsers {
	return &models.TrojanUsers{
		ID:       tu.ID,
		Username: tu.Username,
		Quota:    tu.Quota,
		Download: tu.Download,
		Upload:   tu.Upload,
	}
}
