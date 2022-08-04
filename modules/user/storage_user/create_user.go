package storageuser

import (
	"context"
	"task1/common"
	usermodel "task1/modules/user/model_user"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {
	db := s.db

	db.SetupJoinTable(&usermodel.UserCreate{}, "Roles", &usermodel.UserRole{})

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	//if err := db.Create(data).Error; err != nil {
	//	return common.ErrDB(err)
	//}
	//
	//var arrUserRole []usermodel.UserRole
	//
	//for _, value := range data.Roles {
	//	arrUserRole = append(arrUserRole, usermodel.UserRole{
	//		UserId:    data.Id,
	//		Role_code: value.Role_code,
	//	})
	//}
	//
	//if err := db.Create(&arrUserRole).Error; err != nil {
	//	return common.ErrDB(err)
	//}

	return nil
}
