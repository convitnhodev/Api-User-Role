package model_role

type Permission struct {
	PermissionCode string `json:"permission_code" gorm:"column:permission_code"`
	NamePermission string `json:"name_permission" gorm:"column:name_permission"`
}

func (rp Permission) TableName() string {
	return "permissions"
}
