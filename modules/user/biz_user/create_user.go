package bizuser

import (
	"context"
	"task1/common"
	usermodel "task1/modules/user/model_user"
)

type CreateUserStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type createUserBiz struct {
	store  CreateUserStore
	hasher Hasher
}

func NewCreateUserBiz(store CreateUserStore, hasher Hasher) *createUserBiz {
	return &createUserBiz{store, hasher}
}

func (biz *createUserBiz) CreateNewUser(ctx context.Context, data *usermodel.UserCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	user, err := biz.store.FindUser(ctx, map[string]interface{}{"email": data.Email}, "Roles")
	if user != nil {
		return common.ErrEntityExisted("User Register", err)
	}

	salt := common.GenSalt(50)

	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt

	if err := biz.store.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity("User Register", err)
	}
	return nil
}
