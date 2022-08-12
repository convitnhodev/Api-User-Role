package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"strings"
	"task1/component"
	"task1/component/tokenprovider"
	"time"
)

type jwtProvider struct {
	secret string
}

func NewTokenJWTProvider(secret string) *jwtProvider {
	return &jwtProvider{
		secret: secret,
	}
}

type myClaims struct {
	Payload tokenprovider.TokenPayload `json:"payload"`
	jwt.StandardClaims
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayload, expiry float32) (*tokenprovider.Token, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		Payload: data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(expiry) * time.Second).Unix(),
			IssuedAt:  time.Now().UTC().Unix(),
		},
	})

	myToken, err := t.SignedString([]byte(j.secret))

	if err != nil {
		return nil, err
	}

	// return the token
	return &tokenprovider.Token{
		Token:   myToken,
		Expiry:  expiry,
		Created: time.Now(),
	}, nil
}

func (j *jwtProvider) Validate(token string, ctx component.AppContext, flag *bool) (*tokenprovider.TokenPayload, error) {
	*flag = false
	res, err := jwt.ParseWithClaims(token, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil && !strings.Contains(err.Error(), "token is expired") {
		return nil, tokenprovider.ErrInvalidToken
	}

	if err != nil && strings.Contains(err.Error(), "token is expired") {
		if res.Claims.(*myClaims).ExpiresAt-res.Claims.(*myClaims).IssuedAt == int64(ctx.GetTimeJWT().TimeAccess) {
			return nil, tokenprovider.ErrInvalidToken1
		}
		return nil, tokenprovider.ErrInvalidToken2
	}

	if res.Claims.(*myClaims).ExpiresAt-res.Claims.(*myClaims).IssuedAt > int64(ctx.GetTimeJWT().TimeAccess) {
		*flag = true
	}

	if !res.Valid {
		return nil, tokenprovider.ErrInvalidToken
	}

	claims, ok := res.Claims.(*myClaims)
	if !ok {
		return nil, tokenprovider.ErrInvalidToken
	}
	return &claims.Payload, nil
}
