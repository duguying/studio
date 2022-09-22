package db

import (
	"duguying/studio/modules/dbmodels"

	"gorm.io/gorm"
)

// ListCover 列出封面
func ListCover(tx *gorm.DB) (covers []*dbmodels.Cover, err error) {
	err = tx.Model(dbmodels.Cover{}).Order("created_at desc").Find(&covers).Error
	if err != nil {
		return nil, err
	}
	return covers, nil
}
