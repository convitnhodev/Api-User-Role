package storageuser

import (
	"context"
	"task1/common"
	usermodel "task1/modules/user/model_user"
)

func (s *sqlStore) DeleteUser(ctx context.Context, conditions map[string]interface{}) error {
	db := s.db

	if err := db.Table(usermodel.User{}.TableName()).Where(conditions).Update("active", 0).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
