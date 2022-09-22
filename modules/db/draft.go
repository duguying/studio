package db

import (
	"duguying/studio/modules/dbmodels"

	"gorm.io/gorm"
)

// AddDraft 添加草稿
func AddDraft(tx *gorm.DB, articleID uint, content string) error {
	return tx.Model(dbmodels.Draft{}).Create(&dbmodels.Draft{
		ArticleID: articleID,
		Content:   content,
	}).Error
}
