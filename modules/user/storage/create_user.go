package storageuser

import (
	 "context"
	 usermodel "task1/modules/user/model"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {
	 db := s.db.Begin()

	 // create new user
	 if err := db.Create(data).Error; err != nil {
		  db.Rollback()
		  return err
	 }

	 if err := db.Commit().Error; err != nil {
		  return err
	 }
	 return nil
}
