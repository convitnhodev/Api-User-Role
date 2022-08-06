package biz_role

import (
	"context"
	"task1/common"
	"task1/modules/role/model_role"
)

type ListRoleStore interface {
	ListRoleByConditions(ctx context.Context,
		filter *model_role.Filter,
		conditions map[string]interface{},
		paging *common.Paging,
		moreKeys ...string) ([]model_role.Role, error)
}

type listRoleBiz struct {
	store ListRoleStore
}

func NewListRoleBiz(store ListRoleStore) *listRoleBiz {
	return &listRoleBiz{store}
}

func (biz *listRoleBiz) ListRole(ctx context.Context, filter *model_role.Filter, paging *common.Paging) ([]model_role.Role, error) {

	result, err := biz.store.ListRoleByConditions(ctx, nil, nil, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity("User", err)
	}

	return result, nil
}
