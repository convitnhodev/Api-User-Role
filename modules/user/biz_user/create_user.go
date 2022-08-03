package bizuser

import (
	"context"
	"task1/common"
	usermodel "task1/modules/user/model_user"
)

type CreateUserStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
	UpdateUser(ctx context.Context, data *usermodel.UserUpdate, conditions map[string]interface{}) error
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
	if err := data.Validata(); err != nil {
		return err
	}
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"email": data.Email})
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
