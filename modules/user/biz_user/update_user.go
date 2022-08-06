package bizuser

import (
	"context"
	"errors"
	"task1/common"
	usermodel "task1/modules/user/model_user"
)

type UpdateUserStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*usermodel.User, error)
	UpdateUser(ctx context.Context,
		data *usermodel.UserUpdate,
		conditions map[string]interface{}) error
}

type updateUserBiz struct {
	store  UpdateUserStore
	hasher Hasher
}

func NewUpdateUserBiz(store UpdateUserStore, hasher Hasher) *updateUserBiz {
	return &updateUserBiz{store, hasher}
}

func (biz *updateUserBiz) UpdateUser(ctx context.Context, id int, data *usermodel.UserUpdate) error {

	oldUser, err := biz.store.FindUser(ctx, map[string]interface{}{"user_id": id})
	if err != nil {
		return err
	}

	if oldUser.Active == 0 {
		return errors.New("data deleted")
	}

	if biz.hasher.Hash(data.Password+oldUser.Salt) != oldUser.Password {
		salt := common.GenSalt(50)
		data.Password = biz.hasher.Hash(data.Password + salt)
		data.Salt = salt
	} else {
		data.Password = oldUser.Password
	}

	if err := biz.store.UpdateUser(ctx, data, map[string]interface{}{"user_id": id}); err != nil {
		return err
	}
	return nil
}
