package modelrole

import (
	"task1/common"
	moduledept "task1/modules/dept/moduleDept"
	modulepermission "task1/modules/permission/modulePermission"
)

type Role struct {
	common.SQLModel `json:",inline"`
	RoleCode        string                         `json:"role_code" gorm:"column:role_code;primary_key"`
	RoleName        string                         `json:"name_role" gorm:"column:name_role"`
	DepartmentId    string                         `json:"department_id" gorm:"column:dept_id"`
	Dept            moduledept.Dept                `gorm:"foreignKey:DepartmentId;references:DeptId"`
	Permissions     []*modulepermission.Permission `json:"permissions" gorm:"many2many:role_permission;ForeignKey:role_code;joinForeignKey:role_code;References:permission_code;joinReferences:permission_code"`
}

func (r Role) TableName() string {
	return "roles"
}
