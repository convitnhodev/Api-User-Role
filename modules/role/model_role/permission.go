package model_role

type Permission struct {
	Permission_code string `json:"permission_code" gorm:"column:permission_code"`
	Name_permission string `json:"name_permission" gorm:"column:name_permission"`
}
