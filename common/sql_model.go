package common

import (
	"time"
)

type SQLModel struct {
	CreateAt *time.Time `json:"created_at,omitempty" gorm:"-"`
	UpdateAt *time.Time `json:"updated_at,omitempty" gorm:"-"`
}
