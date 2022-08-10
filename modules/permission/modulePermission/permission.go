package modulepermission

type Permission struct {
	PermissionCode string `json:"permission_code" gorm:"column:permission_code;primaryKey"`
	NamePermission string `json:"name_permission" gorm:"column:name_permission"`
}

func (rp Permission) TableName() string {
	return "permissions"
}
