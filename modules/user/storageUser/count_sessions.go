package storageuser

import (
	"context"
	"task1/common"
	usermodel "task1/modules/user/modelUser"
	"task1/modules/userControl/modelUserControl"
	"time"
)

const timeLayout = "2006-01-02 15:04:05"

func (s *sqlStore) CountSessions(ctx context.Context, timeBegin *time.Time, timeEnd *time.Time, email string) (*int, error) {
	db := s.db

	var data usermodel.SqlData

	if err := db.Select("email, count(email) as count").
		Table(modelUserControl.Session{}.TableName()).
		Where("email = ?", email).
		Where("created_at >= ?", timeBegin.Format(timeLayout)).
		Where("created_at <= ?", timeEnd.Format(timeLayout)).
		Group("email").
		First(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &data.Count, nil
}

func (s *sqlStore) ListSessions(ctx context.Context, filter *usermodel.Filter, timeBegin *time.Time, timeEnd *time.Time, email []string) ([]usermodel.SqlData, error) {
	db := s.db

	var listSessions []usermodel.SqlData

	if err := db.Select("email, count(email) as count").
		Table(modelUserControl.Session{}.TableName()).
		Where(filter).
		Where("email in (?)", email).
		Where("created_at >= ?", timeBegin.Format(timeLayout)).
		Where("created_at <= ?", timeEnd.Format(timeLayout)).
		Group("email").
		Find(&listSessions).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return listSessions, nil
}
