package bizrole

import (
	"context"
	modelrole "task1/modules/role/modelRole"
)

type GetRoleStore interface {
	FindRole(ctx context.Context, conditions map[string]interface{}) (*modelrole.Role, error)
}

type getRoleBiz struct {
	store GetRoleStore
}

func NewGetRoleBiz(store GetRoleStore) *getRoleBiz {
	return &getRoleBiz{store}
}

func (biz *getRoleBiz) GetRole(ctx context.Context, conditions map[string]interface{}) (*modelrole.Role, error) {

	role, err := biz.store.FindRole(ctx, conditions)
	if err != nil {
		return nil, err
	}

	return role, nil
}
