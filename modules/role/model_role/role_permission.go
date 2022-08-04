package model_role

type RolePermission struct {
	Role_code       string `json:"role_code" gorm:"column:role_code"`
	Permission_code string `json:"permission_code" gorm:"column:permission_code"`
}

func (rp RolePermission) TableName() string {
	return "role_permission"
}
