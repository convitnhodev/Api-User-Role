package storagerole

import (
	"context"
	"gorm.io/gorm/clause"
	"task1/common"
	"task1/modules/role/model_role"
)

func (s *sqlStore) ListRoleByConditions(ctx context.Context,
	filter *model_role.Filter,
	conditions map[string]interface{},
	paging *common.Paging,
	moreKeys ...string) ([]model_role.Role, error) {
	db := s.db
	var data []model_role.Role

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Table(model_role.Role{}.TableName()).Count(&paging.Total).Error; err != nil {
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
