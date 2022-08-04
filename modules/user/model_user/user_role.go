package usermodel

type UserRole struct {
	UserId    int    `json:"user_id" gorm:"column:user_id;primary_key"`
	Role_code string `json:"role_code" gorm:"column:role_code;primary_key"`
}

func (ur UserRole) TableName() string {
	return "user_role"
}
