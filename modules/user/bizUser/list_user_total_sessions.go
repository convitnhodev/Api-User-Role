package bizuser

import (
	"context"
	"task1/common"
	usermodel "task1/modules/user/modelUser"

	"time"
)

type UserListLoginStore interface {
	ListUserByConditions(ctx context.Context,
		filter *usermodel.Filter,
		conditions map[string]interface{},
		paging *common.Paging,
		moreKeys ...string) ([]usermodel.User, error)
	ListSessions(ctx context.Context, filter *usermodel.Filter, timeBegin *time.Time, timeEnd *time.Time, email []string) ([]usermodel.SqlData, error)
}

type userListLoginBiz struct {
	store UserListLoginStore
}

func NewUserListLoginBiz(store UserListLoginStore) *userListLoginBiz {
	return &userListLoginBiz{store}
}

func (biz *userListLoginBiz) ListSessionsWeek(ctx context.Context, filter *usermodel.Filter, paging *common.Paging) ([]usermodel.SqlData, error) {
	listEmail, err := biz.store.ListUserByConditions(ctx, filter, nil, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity("User", err)
	}
	arr := make([]string, 0)
	for _, item := range listEmail {
		arr = append(arr, item.Email)
	}
	TimeNow := time.Now()
	LastTime := TimeNow.AddDate(0, 0, -7)
	result, err := biz.store.ListSessions(ctx, nil, &LastTime, &TimeNow, arr)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (biz *userListLoginBiz) ListSessionsMonth(ctx context.Context, filter *usermodel.Filter, paging *common.Paging) ([]usermodel.SqlData, error) {
	listEmail, err := biz.store.ListUserByConditions(ctx, filter, nil, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity("User", err)
	}
	arr := make([]string, 0)
	for _, item := range listEmail {
		arr = append(arr, item.Email)
	}
	TimeNow := time.Now()
	LastTime := TimeNow.AddDate(0, -1, 0)
	result, err := biz.store.ListSessions(ctx, nil, &LastTime, &TimeNow, arr)
	if err != nil {
		return nil, err
	}
	return result, nil
}
