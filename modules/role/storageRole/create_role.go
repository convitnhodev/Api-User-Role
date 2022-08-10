package storagerole

import (
	"context"
	"task1/common"
	modelrole "task1/modules/role/modelRole"
)

func (s *sqlStore) CreateRole(ctx context.Context, data *modelrole.Role) error {
	db := s.db

	// create new role
	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
