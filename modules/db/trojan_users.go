package db

import (
	"duguying/studio/modules/dbmodels"

	"gorm.io/gorm"
)

// ListAllTrojanUsers 列举所有trojan帐号
func ListAllTrojanUsers(tx *gorm.DB) (list []*dbmodels.TrojanUsers, err error) {
	list = []*dbmodels.TrojanUsers{}
	err = tx.Model(dbmodels.TrojanUsers{}).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}
