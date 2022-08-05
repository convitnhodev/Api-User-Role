package model_role

type RolePermission struct {
	RoleCode       string `json:"role_code" gorm:"column:role_code"`
	PermissionCode string `json:"permission_code" gorm:"column:permission_code"`
}

func (rp RolePermission) TableName() string {
	return "role_permission"
}
