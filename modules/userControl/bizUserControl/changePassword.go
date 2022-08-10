package bizUserControl

import (
	"context"
	"errors"
	"task1/common"
	"task1/modules/userControl/modelUserControl"
)

type ChangePasswordStore interface {
	UpdatePassword(ctx context.Context, conditions map[string]interface{}, newPassword string, newSalt string) error
}

type changePasswordBiz struct {
	store  ChangePasswordStore
	hasher Hasher
}

func NewChangPasswordBiz(store ChangePasswordStore, hasher Hasher) *changePasswordBiz {
	return &changePasswordBiz{store, hasher}
}

func (biz *changePasswordBiz) ChangePassword(ctx context.Context, data common.Requester, oldPassword string, newPassword modelUserControl.PASSWORD) error {
	if data.GetPassword() != biz.hasher.Hash(oldPassword+data.GetSalt()) {
		return common.ErrInvalidPassword(errors.New("invalid password"))
	}

	if err := newPassword.Validate(); err != nil {
		return err
	}

	newSalt := common.GenSalt(50)

	biz.store.UpdatePassword(ctx, map[string]interface{}{"user_id": data.GetUserId()}, biz.hasher.Hash(string(newPassword)+newSalt), newSalt)
	return nil
}
