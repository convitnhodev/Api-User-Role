package bizuser

import (
	"context"
	usermodel "task1/modules/user/modelUser"
)

type GetUserStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*usermodel.User, error)
}

type getUserBiz struct {
	store GetUserStore
}

func NewGetUserBiz(store GetUserStore) *getUserBiz {
	return &getUserBiz{store}
}

func (biz *getUserBiz) GetUser(ctx context.Context, id int) (*usermodel.User, error) {

	user, err := biz.store.FindUser(ctx, map[string]interface{}{"user_id": id}, "")
	if err != nil {
		return nil, err
	}

	return user, nil
}
