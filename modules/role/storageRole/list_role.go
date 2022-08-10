package storagerole

import (
	"context"
	"gorm.io/gorm/clause"
	"task1/common"
	modelrole "task1/modules/role/modelRole"
)

func (s *sqlStore) ListRoleByConditions(ctx context.Context,
	filter *modelrole.Filter,
	conditions map[string]interface{},
	paging *common.Paging,
	moreKeys ...string) ([]modelrole.Role, error) {
	db := s.db
	var data []modelrole.Role

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Table(modelrole.Role{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := db.Where(conditions).
		Where(filter).
		Limit(paging.Limit).
		Offset((paging.Page - 1) * paging.Limit).
		Preload(clause.Associations).
		Find(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return data, nil
}
