package bizUserControl

import (
	 "context"
	 usermodel "task1/modules/user/modelUser"
)

type ChangePasswordStore interface {
	 UpdatePassword(ctx context.Context, conditions map[string]interface{}, newPassword string, newSalt string) (*usermodel.User, error)
}

type changePasswordBiz struct {
	 store  ChangePasswordStore
	 hasher Hasher
}

func NewChangPasswordBiz(store ChangePasswordStore, hasher Hasher) *changePasswordBiz {
	 return &changePasswordBiz{store, hasher}
}

func (biz *changePasswordBiz) ChangePassword(ctx context.Context, data) error {

}
