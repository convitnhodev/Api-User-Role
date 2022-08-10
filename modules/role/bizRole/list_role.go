package bizrole

import (
	"context"
	"task1/common"
	modelrole "task1/modules/role/modelRole"
)

type ListRoleStore interface {
	ListRoleByConditions(ctx context.Context,
		filter *modelrole.Filter,
		conditions map[string]interface{},
		paging *common.Paging,
		moreKeys ...string) ([]modelrole.Role, error)
}

type listRoleBiz struct {
	store ListRoleStore
}

func NewListRoleBiz(store ListRoleStore) *listRoleBiz {
	return &listRoleBiz{store}
}

func (biz *listRoleBiz) ListRole(ctx context.Context, filter *modelrole.Filter, paging *common.Paging) ([]modelrole.Role, error) {

	result, err := biz.store.ListRoleByConditions(ctx, nil, nil, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity("User", err)
	}

	return result, nil
}
