package common

import (
	"time"
)

type SQLModel struct {
	CreateAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdateAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}
