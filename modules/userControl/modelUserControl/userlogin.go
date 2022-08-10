package modelUserControl

import (
	"errors"
	"task1/common"
)

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email" form:"email"`
	Password string `json:"password" form:"password" gorm:"column:password form:password"`
}

func (u UserLogin) TableName() string {
	return "users"
}

type PASSWORD string

func (pass PASSWORD) Validate() error {
	if len(pass) < 6 {
		return common.ErrInvalidPassword(errors.New("password must be at least 6 characters"))
	}

	return nil

}
