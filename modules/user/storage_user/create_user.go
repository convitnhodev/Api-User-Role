package storageuser

import (
	"context"
	"task1/common"
	usermodel "task1/modules/user/model_user"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {
	db := s.db.Begin()

	// create new user
	if err := db.Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
