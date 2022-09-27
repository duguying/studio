// Package db 包注释
package db

import (
	"duguying/studio/modules/dbmodels"
	"time"

	"gorm.io/gorm"
)

// AddCalendar 添加日志事件
func AddCalendar(tx *gorm.DB, date time.Time,
	summary, address, description, link, attendee string) (added *dbmodels.Calendar, err error) {
	added = &dbmodels.Calendar{
		Date:        date,
		Summary:     summary,
		Address:     address,
		Description: description,
		Link:        link,
		Attendee:    attendee,
	}
	err = tx.Model(dbmodels.Calendar{}).Create(added).Error
	if err != nil {
		return nil, err
	}
	return added, nil
}

// ListAllCalendarIds 列举所有未发送的日历事件表
func ListAllCalendarIds(tx *gorm.DB) (ids []string, err error) {
	list := []*dbmodels.Calendar{}
	err = tx.Model(dbmodels.Calendar{}).Select("id").Where("send_at is NULL").Find(&list).Error
	if err != nil {
		return nil, err
	}
	ids = []string{}
	for _, calendar := range list {
		ids = append(ids, calendar.ID)
	}
	return ids, nil
}

// GetCalendarByID 按ID获取日历
func GetCalendarByID(tx *gorm.DB, id string) (calendar *dbmodels.Calendar, err error) {
	calendar = &dbmodels.Calendar{}
	err = tx.Model(dbmodels.Calendar{}).Where("id=?", id).First(calendar).Error
	if err != nil {
		return nil, err
	}
	return calendar, nil
}
