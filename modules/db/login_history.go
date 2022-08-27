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

// PageLoginHistoryByUserID 按用户列举登陆历史
func PageLoginHistoryByUserID(tx *gorm.DB, userID uint, page uint, pageSize uint) (list []*dbmodels.LoginHistory, total int64, err error) {
	total = 0
	list = []*dbmodels.LoginHistory{}

	err = tx.Model(dbmodels.LoginHistory{}).Where("user_id=?", userID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = g.Db.Model(dbmodels.LoginHistory{}).Where("user_id=?", userID).Order("login_at desc").
		Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Find(&list).Error
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
