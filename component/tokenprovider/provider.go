package tokenprovider

import (
	"errors"
	"task1/common"
	"task1/component"
	"time"
)

type Provider interface {
	Generate(data TokenPayload, expiry float32) (*Token, error)
	Validate(token string, ctx component.AppContext, flag *bool) (*TokenPayload, error)
}

var (
	ErrNotFound = common.NewCustomError(
		errors.New("token not found"),
		"token not found",
		"ErrNotFound",
	)

	ErrEncodingToken = common.NewCustomError(
		errors.New("error encoding the token"),
		"error encoding the token",
		"ErrEncodingToken",
	)

	ErrInvalidToken = common.NewCustomError(
		errors.New("invalid token provider"),
		"invalid token provider",
		"ErrInvalidToken",
	)
	ErrInvalidToken1 = common.NewCustomError(
		errors.New("invalid token provider"),
		"invalid token provider (access token expired)",
		"ErrInvalidToken1",
	)

	ErrInvalidToken2 = common.NewCustomError(
		errors.New("invalid token provider"),
		"invalid token provider (refresh token expired)",
		"ErrInvalidToken2",
	)
)

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  float32   `json:"expiry"`
}

type Account struct {
	AccessToken  *Token
	RefreshToken *Token
}

type TokenPayload struct {
	UserId int `json:"user_id"`
}
