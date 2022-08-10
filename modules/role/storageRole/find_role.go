package storagerole

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"task1/common"
	modelrole "task1/modules/role/modelRole"
)

func (s *sqlStore) FindRole(ctx context.Context, conditions map[string]interface{}) (*modelrole.Role, error) {
	db := s.db

	var data modelrole.Role
	// find user,
	if err := db.Table(data.TableName()).Preload(clause.Associations).First(&data, conditions).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
