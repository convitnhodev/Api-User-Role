package storageuser

import (
	"context"
	"task1/common"
	usermodel "task1/modules/user/modelUser"
)

func (s *sqlStore) UpdateUser(ctx context.Context,
	data *usermodel.UserUpdate,
	conditions map[string]interface{}) error {
	db := s.db.Begin()
	if err := db.Table("user_role").Delete(nil, conditions).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Where(conditions).Updates(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
