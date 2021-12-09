package db

import (
	"duguying/studio/modules/dbmodels"
	"duguying/studio/service/models"

	"gorm.io/gorm"
)

func PutApiLog(tx *gorm.DB, apiLog *models.ApiLog) error {
	err := tx.Model(dbmodels.ApiLog{}).Create(apiLog).Error
	if err != nil {
		return err
	}
	return nil
}
