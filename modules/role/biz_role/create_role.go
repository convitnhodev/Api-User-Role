package biz_role

import (
	"context"
	"task1/common"
	"task1/modules/role/model_role"
)

type CreateRoleStore interface {
	CreateRole(ctx context.Context, data *model_role.Role) error
	FindRole(ctx context.Context, conditions map[string]interface{}) (*model_role.Role, error)
}

type CreateRoleBiz struct {
	store CreateRoleStore
}

func NewCreateRoleBiz(store CreateRoleStore) *CreateRoleBiz {
	return &CreateRoleBiz{store}
}

func (biz *CreateRoleBiz) CreateRole(ctx context.Context, data *model_role.Role) error {
	role, err := biz.store.FindRole(ctx, map[string]interface{}{"role_code": data.Role_code})
	if role != nil {
		return common.ErrEntityExisted("User Register", err)
	}
	if err := biz.store.CreateRole(ctx, data); err != nil {
		return common.ErrCannotCreateEntity("Role", err)
	}
	return nil
}
