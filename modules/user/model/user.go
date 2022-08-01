package usermodel

import (
	"task1/common"
	"time"
)

type User struct {
	common.SQLModel `json:",inline"`
	Id              int    `json:"id" gorm:"column:user_id"`
	Email           string `json:"email" gorm:"column:email"`
	Password        string `json:"password" gorm:"column:password"`
	Role            int    `json:"-" gorm:"column:role"`
	Salt            string `json:"-" gorm:"column:salt"`
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Id              int        `json:"id" gorm:"column:user_id"`
	Email           string     `json:"email" gorm:"column:email"`
	LastName        string     `json:"last_name" gorm:"column:last_name"`
	FirstName       string     `json:"first_name" gorm:"column:first_name"`
	Password        string     `json:"password" gorm:"column:password"`
	Role            int        `json:"-" gorm:"column:role"`
	Salt            string     `json:"-" gorm:"column:salt"`
	CreateAt        *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
}

func (u UserCreate) TableName() string {
	return "users"
}

func (u User) TableName() string {
	return "users"
}
