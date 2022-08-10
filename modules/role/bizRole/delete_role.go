package bizrole

import (
	"context"
	"task1/common"
)

type DeleteRoleStore interface {
	DeleteRole(ctx context.Context, conditions map[string]interface{}) error
}

type delteRoleBiz struct {
	store DeleteRoleStore
}

func NewDeleteRoleBiz(store DeleteRoleStore) *delteRoleBiz {
	return &delteRoleBiz{store}
}

func (biz *delteRoleBiz) DeleteRole(ctx context.Context, conditions map[string]interface{}) error {

	if err := biz.store.DeleteRole(ctx, conditions); err != nil {
		return common.ErrEntityDeleted("Role", err)
	}
	return nil
}
