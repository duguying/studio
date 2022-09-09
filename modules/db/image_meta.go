package db

import (
	"duguying/studio/modules/dbmodels"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// AddImageMeta 添加图片 meta 信息
func AddImageMeta(tx *gorm.DB, fileID string, meta, metas string) (added *dbmodels.ImageMeta, err error) {
	added = &dbmodels.ImageMeta{
		FileID: fileID,
		Meta:   datatypes.JSON(meta),
		Metas:  datatypes.JSON(metas),
	}
	err = tx.Model(dbmodels.ImageMeta{}).Create(added).Error
	if err != nil {
		return nil, err
	}
	return added, nil
}
