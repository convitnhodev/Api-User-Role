package storageUserControl

import (
	"context"
	"task1/common"
	"task1/modules/userControl/modelUserControl"
)

func (s *sqlStore) CreateSession(ctx context.Context, email string) error {
	db := s.db

	var data modelUserControl.Session
	data.Email = email
	if err := db.Table(data.TableName()).Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
