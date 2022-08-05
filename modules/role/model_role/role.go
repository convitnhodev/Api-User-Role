package model_role

import (
	"task1/common"
)

type Role struct {
	common.SQLModel `json:",inline"`
	RoleCode        string `json:"role_code" gorm:"column:role_code;primary_key"`
	RoleName        string `json:"role_name" gorm:"column:name_role"`
	DepartmentId    int    `json:"department_id" gorm:"column:dept_id"`
	//Permissions     []Permission `json:"permission" gorm:"many2many:role_permission;"`
}

func (r Role) TableName() string {
	return "roles"
}
