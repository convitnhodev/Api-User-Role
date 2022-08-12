package modelUserControl

import (
	"task1/component/tokenprovider"
	"time"
)

type Session struct {
	Id       int        `json:"id" gorm:"column:id_session;primaryKey;default:0"`
	Email    string     `json:"email" form:"email" gorm:"column:email" form:"email"`
	CreateAt *time.Time `json:"created_at,omitempty" gorm:"-"`
}

type Survivor struct {
	Email    string                `json:"email" form:"email" gorm:"column:email" form:"email"`
	Account  tokenprovider.Account `json:"account" form:"account" gorm:"column:account" form:"account"`
	IdDevice string                `json:"id_device" form:"id_device" gorm:"column:id_device" form:"id_device"`
}

func (u Survivor) TableName() string {
	return "survivors"
}

func (u Session) TableName() string {
	return "sessions"
}
