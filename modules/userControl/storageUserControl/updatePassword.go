package storageUserControl

import (
	"context"
	"fmt"
	"task1/common"
	"task1/modules/userControl/modelUserControl"
)

func (s *sqlStore) UpdatePassword(ctx context.Context, conditions map[string]interface{}, newPassword string, newSalt string) error {
	db := s.db

	fmt.Println("hello")

	if err := db.Table(modelUserControl.UserLogin{}.TableName()).Where(conditions).Updates(map[string]interface{}{"salt": newSalt, "password": newPassword}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
