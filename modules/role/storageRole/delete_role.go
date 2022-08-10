package storagerole

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"task1/common"
	modelrole "task1/modules/role/modelRole"
)

func (s *sqlStore) DeleteRole(ctx context.Context, conditions map[string]interface{}) error {
	//db := s.db.Begin()
	//
	//if err := db.Table(model_role.Role{}.TableName()).Where(conditions).Delete(nil).Error; err != nil {
	//	db.Rollback()
	//	return common.ErrDB(err)
	//}
	//
	//if err := db.Commit().Error; err != nil {
	//	return common.ErrDB(err)
	//}
	//
	db := s.db

	var role *modelrole.Role

	if err := db.Preload("Permissions").Find(&role, conditions).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.RecordNotFound
		}
		return common.ErrDB(err)
	}

	//db.Model(&role).Delete(&role).Association("Permissions").Delete(&role.Permissions)
	if err := db.Select(clause.Associations).Delete(&role).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
