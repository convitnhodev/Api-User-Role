package storageuser

import (
	"context"
	"gorm.io/gorm"
	"task1/common"
	usermodel "task1/modules/user/model_user"
)

func (s *sqlStore) UpdateUser(ctx context.Context,
	data *usermodel.UserUpdate,
	conditions map[string]interface{}) error {
	db := s.db

	if err := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Where(conditions).Table("users").
		Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
