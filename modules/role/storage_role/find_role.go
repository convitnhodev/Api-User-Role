package storagerole

import (
	"context"
	"gorm.io/gorm"
	"task1/common"
	"task1/modules/role/model_role"
)

func (s *sqlStore) FindRole(ctx context.Context, conditions map[string]interface{}) (*model_role.Role, error) {
	db := s.db

	var data model_role.Role
	// find user,
	if err := db.Table(data.TableName()).First(&data, conditions).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
