package storageuser

import (
	"context"
	"task1/common"
	usermodel "task1/modules/user/model_user"
)

func (s *sqlStore) ListUserByConditions(ctx context.Context,
	filter *usermodel.Filter,
	conditions map[string]interface{},
	paging *common.Paging,
	moreKeys ...string) ([]usermodel.User, error) {
	db := s.db
	var data []usermodel.User

	db = db.Where("active in (1)")

	if err := db.Table(usermodel.User{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(conditions).
		Where(filter).
		Limit(paging.Limit).
		Offset((paging.Page - 1) * paging.Limit).
		Find(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return data, nil
}
