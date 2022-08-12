package component

import (
	"gorm.io/gorm"
)

type TimeJWT struct {
	TimeAccess  float32
	TimeRefresh float32
}

type AppContext interface {
	GetMainDbConnection() *gorm.DB
	SecretKey() string
	GetTimeJWT() TimeJWT
}

type appCtx struct {
	db        *gorm.DB
	secretKey string
	TimeJWT   TimeJWT
}

func NewAppContext(db *gorm.DB, secretkey string, timeJWT TimeJWT) *appCtx {
	return &appCtx{db, secretkey, timeJWT}
}

func (ctx *appCtx) GetMainDbConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secretKey
}

func (ctx *appCtx) GetTimeJWT() TimeJWT {
	return ctx.TimeJWT
}
