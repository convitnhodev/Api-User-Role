package bizuser

import (
	"context"

	"time"
)

type UserLoginStore interface {
	CountSessions(ctx context.Context, timeBegin *time.Time, timeEnd *time.Time, email string) (*int, error)
}

type userLoginBiz struct {
	store UserLoginStore
}

func NewUserLoginWeekBiz(store UserLoginStore) *userLoginBiz {
	return &userLoginBiz{store}
}

func (biz *userLoginBiz) CountSessionsWeek(ctx context.Context, email string) (*int, error) {
	TimeNow := time.Now()
	LastTime := TimeNow.AddDate(0, 0, -7)
	total, err := biz.store.CountSessions(ctx, &LastTime, &TimeNow, email)
	if err != nil {
		return nil, err
	}
	return total, nil
}

func (biz *userLoginBiz) CountSessionsMonth(ctx context.Context, email string) (*int, error) {
	TimeNow := time.Now()
	LastTime := TimeNow.AddDate(0, -1, 0)
	total, err := biz.store.CountSessions(ctx, &LastTime, &TimeNow, email)
	if err != nil {
		return nil, err
	}
	return total, nil
}
