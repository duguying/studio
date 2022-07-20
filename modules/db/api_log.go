package db

import (
	"duguying/studio/modules/dbmodels"
	"duguying/studio/service/models"

	"gorm.io/gorm"
)

// PutApiLog db中记录日志
func PutApiLog(tx *gorm.DB, apiLog *models.APILog) error {
	err := tx.Model(dbmodels.APILog{}).Create(apiLog).Error
	if err != nil {
		return err
	}
	return nil
}
