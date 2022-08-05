package usermodel

import (
	"errors"
	"strings"
	"task1/common"
	"task1/modules/role/model_role"
)

type User struct {
	Id              int `json:"id" gorm:"column:user_id;primary_key"`
	Active          int `json:"active" gorm:"column:active;default:1"`
	common.SQLModel `json:",inline"`
	Email           string            `json:"email" gorm:"column:email"`
	Password        string            `json:"password" gorm:"column:password"`
	Salt            string            `json:"-" gorm:"column:salt"`
	Roles           []model_role.Role `json:"roles" gorm:"many2many:user_role;"`
}

type UserCreate struct {
	Id              int `json:"id" gorm:"column:user_id;primaryKey;autoIncrement"`
	Active          int `json:"active" gorm:"column:active;default:1"`
	common.SQLModel `json:",inline"`
	Email           string            `json:"email" gorm:"column:email"`
	LastName        string            `json:"last_name" gorm:"column:last_name"`
	FirstName       string            `json:"first_name" gorm:"column:first_name"`
	Password        string            `json:"password" gorm:"column:password"`
	Salt            string            `json:"-" gorm:"column:salt"`
	Roles           []model_role.Role `json:"roles" gorm:"many2many:user_role;ForeignKey:user_id;joinForeignKey:user_id;References:role_code;joinReferences:role_code"`
}

type UserUpdate struct {
	Id              int `json:"id" gorm:"column:user_id"`
	Active          int `json:"active" gorm:"column:active;default:1"`
	common.SQLModel `json:",inline"`
	Email           string            `json:"email" gorm:"column:email"`
	LastName        *string           `json:"last_name" gorm:"column:last_name"`
	FirstName       *string           `json:"first_name" gorm:"column:first_name"`
	Password        string            `json:"password" gorm:"column:password"`
	Role            []model_role.Role `json:"role" gorm:"many2many:user_role;"`
	Salt            string            `json:"-" gorm:"column:salt"`
}

func (u UserCreate) TableName() string {
	return "users"
}

func (u UserUpdate) TableName() string {
	return "users"
}

func (u User) TableName() string {
	return "users"
}

func (user *UserCreate) Validate() error {

	// check validate of email
	user.Email = strings.TrimSpace(user.Email)
	user.LastName = strings.TrimSpace(user.LastName)
	user.FirstName = strings.TrimSpace(user.LastName)
	user.Password = strings.TrimSpace(user.Password)

	if user.Email == "" {
		return errors.New("email name can not be blank")
	}

	if user.Password == "" {
		return errors.New("password name can not be blank")
	}

	if user.FirstName == "" || user.LastName == "" {
		return errors.New("firstname or lastname name can not be blank")
	}

	return nil
}
