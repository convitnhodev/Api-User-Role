package bizuser

import (
	"context"
	"task1/common"
	usermodel "task1/modules/user/modelUser"
)

type ListUserStore interface {
	ListUserByConditions(ctx context.Context,
		filter *usermodel.Filter,
		conditions map[string]interface{},
		paging *common.Paging,
		moreKeys ...string) ([]usermodel.User, error)
}

type listUserBiz struct {
	store ListUserStore
}

func NewListUserBiz(store ListUserStore) *listUserBiz {
	return &listUserBiz{store: store}
}

func (biz *listUserBiz) ListUser(ctx context.Context, filter *usermodel.Filter, paging *common.Paging) ([]usermodel.User, error) {
	result, err := biz.store.ListUserByConditions(ctx, filter, nil, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity("User", err)
	}

	return result, nil
}
