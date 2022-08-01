package common

import (
	"time"
)

type SQLModel struct {
	Id       int        `json:"id" gorm:"column:user_id"`
	Active   int        `json:"active" gorm:"column:active;default:1"`
	CreateAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdateAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}
