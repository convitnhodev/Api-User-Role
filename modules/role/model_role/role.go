package model_role

import (
	"task1/common"
)

type Role struct {
	common.SQLModel `json:",inline"`
	Role_code       int    `json:"role_code" gorm:"column:role_code"`
	Role_name       string `json:"role_name" gorm:"column:role_name"`
	Department_id   int    `json:"department_id" gorm:"column:dept_id"`
	//Permissions     []Permission `json:"permission" gorm:"many2many:role_permission;"`
}

func (r Role) TableName() string {
	return "role"
}
