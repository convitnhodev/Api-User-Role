package storageUserControl

import (
	"context"
	"task1/common"
)

func (s *sqlStore) UpdatePassword(ctx context.Context, conditions map[string]interface{}, newPassword string, newSalt string) error {
	db := s.db

	if err := db.Where(conditions).Update("salt", newSalt).Update("password", newPassword).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
