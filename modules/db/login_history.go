package db

import (
	"duguying/studio/g"
	"duguying/studio/modules/dbmodels"
	"duguying/studio/modules/session"

	"gorm.io/gorm"
)

// AddLoginHistory 添加登陆历史
func AddLoginHistory(tx *gorm.DB, sessionID string, entity *session.Entity) error {
	hist := &dbmodels.LoginHistory{
		UserID:    entity.UserID,
		SessionID: sessionID,
		IP:        entity.IP,
		LoginAt:   &entity.LoginAt,
	}
	return tx.Model(dbmodels.LoginHistory{}).Create(hist).Error
}

// ListLoginHistoryByUserID 按用户列举登陆历史
func ListLoginHistoryByUserID(userID uint) (list []*dbmodels.LoginHistory, err error) {
	list = []*dbmodels.LoginHistory{}
	err = g.Db.Model(dbmodels.LoginHistory{}).Where("user_id=?", userID).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}
