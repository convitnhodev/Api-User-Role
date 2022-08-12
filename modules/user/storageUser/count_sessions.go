package storageuser

import (
	"context"
	"task1/common"
	"task1/modules/userControl/modelUserControl"
	"time"
)

const timeLayout = "2006-01-02 15:04:05"

func (s *sqlStore) CountSessions(ctx context.Context, timeBegin *time.Time, timeEnd *time.Time, email string) (*int, error) {
	db := s.db

	type sqlData struct {
		Email string `gorm:"column:email"`
		Count int    `gorm:"column:count"`
	}

	var data sqlData

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
