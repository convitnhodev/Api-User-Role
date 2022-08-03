package storagerole

import (
	"context"
	"task1/common"
	"task1/modules/role/model_role"
)

func (s *sqlStore) DeleteRole(ctx context.Context, conditions map[string]interface{}) error {
	db := s.db

	if err := db.Table(model_role.Role{}.TableName()).Where(conditions).Delete(conditions).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
