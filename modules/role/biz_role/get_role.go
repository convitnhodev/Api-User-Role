package biz_role

import (
	"context"
	"task1/modules/role/model_role"
)

type GetRoleStore interface {
	FindRole(ctx context.Context, conditions map[string]interface{}) (*model_role.Role, error)
}

type getRoleBiz struct {
	store GetRoleStore
}

func NewGetRoleBiz(store GetRoleStore) *getRoleBiz {
	return &getRoleBiz{store}
}

func (biz *getRoleBiz) GetRole(ctx context.Context, conditions map[string]interface{}) (*model_role.Role, error) {

	role, err := biz.store.FindRole(ctx, conditions)
	if err != nil {
		return nil, err
	}

	return role, nil
}
