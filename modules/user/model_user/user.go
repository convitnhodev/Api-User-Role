package usermodel

import (
	"errors"
	"strings"
	"task1/common"
)

type User struct {
	Id              int `json:"id" gorm:"column:user_id"`
	Active          int `json:"active" gorm:"column:active;default:1"`
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email"`
	Password        string `json:"password" gorm:"column:password"`
	Role            int    `json:"-" gorm:"column:role"`
	Salt            string `json:"-" gorm:"column:salt"`
	//Dept            int    `json:"-" gorm:"column:dept"`
}

type UserCreate struct {
	Id              int `json:"id" gorm:"column:user_id"`
	Active          int `json:"active" gorm:"column:active;default:1"`
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email"`
	LastName        string `json:"last_name" gorm:"column:last_name"`
	FirstName       string `json:"first_name" gorm:"column:first_name"`
	Password        string `json:"password" gorm:"column:password"`
	Role            int    `json:"role" gorm:"column:role"`
	Salt            string `json:"-" gorm:"column:salt"`
	//Dept            int    `json:"-" gorm:"column:dept"`
}

type UserUpdate struct {
	Id              int `json:"id" gorm:"column:user_id"`
	Active          int `json:"active" gorm:"column:active;default:1"`
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email"`
	LastName        string `json:"last_name" gorm:"column:last_name"`
	FirstName       string `json:"first_name" gorm:"column:first_name"`
	Password        string `json:"password" gorm:"column:password"`
	Role            *int   `json:"role" gorm:"column:role"`
	Salt            string `json:"-" gorm:"column:salt"`
	//Dept            int    `json:"-" gorm:"column:dept"`
}

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email" form:"email"`
	Password string `json:"password" form:"password" gorm:"column:password form:password"`
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

func (u UserLogin) TableName() string {
	return "users"
}

func (user *UserCreate) Validata() error {

	// check validate of email
	user.Email = strings.TrimSpace(user.Email)
	if user.Email == "" {
		return errors.New("email name can not be blank")
	}

	// check validate of firstname
	user.FirstName = strings.TrimSpace(user.FirstName)
	if user.FirstName == "" {
		return errors.New("first name can not be blank")
	}

	user.LastName = strings.TrimSpace(user.LastName)
	if user.LastName == "" {
		return errors.New("last name can not be blank")
	}

	return nil
}
