package db

import (
	"duguying/studio/g"
	"duguying/studio/modules/dbmodels"
	"duguying/studio/service/models"
)

func PutApiLog(apiLog *models.ApiLog) error {
	errs := g.Db.Model(dbmodels.ApiLog{}).Create(apiLog).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return errs[0]
	}
	return nil
}
