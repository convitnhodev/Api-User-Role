package bizrole

import (
	"context"
	"task1/common"
	modelrole "task1/modules/role/modelRole"
)

type CreateRoleStore interface {
	CreateRole(ctx context.Context, data *modelrole.Role) error
	FindRole(ctx context.Context, conditions map[string]interface{}) (*modelrole.Role, error)
}

type createRoleBiz struct {
	store CreateRoleStore
}

func NewCreateRoleBiz(store CreateRoleStore) *createRoleBiz {
	return &createRoleBiz{store}
}

func (biz *createRoleBiz) CreateRole(ctx context.Context, data *modelrole.Role) error {
	role, err := biz.store.FindRole(ctx, map[string]interface{}{"role_code": data.RoleCode})
	if role != nil {
		return common.ErrEntityExisted("User", err)
	}
	if err := biz.store.CreateRole(ctx, data); err != nil {
		return common.ErrCannotCreateEntity("Role", err)
	}
	return nil
}
