package db

import (
	"duguying/studio/g"
	"duguying/studio/modules/dbmodels"
	"duguying/studio/service/models"
)

func PutApiLog(apiLog *models.ApiLog) error {
	err := g.Db.Model(dbmodels.ApiLog{}).Create(apiLog).Error
	if err != nil {
		return err
	}
	return nil
}
