package storageuser

import (
	"context"
	"task1/common"
	usermodel "task1/modules/user/modelUser"
)

func (s *sqlStore) DeleteUser(ctx context.Context, conditions map[string]interface{}) error {
	db := s.db.Begin()

	if err := db.Table(usermodel.User{}.TableName()).Where(conditions).Update("active", 0).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Table("user_role").Where(conditions).Delete(nil).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
