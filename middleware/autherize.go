package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"task1/common"
	"task1/component"
	"task1/component/tokenprovider/jwt"
	storageuser "task1/modules/user/storage"
)

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}
	return parts[1], nil
}

// 1. Get token from header
// 2. Validate token and parse to payload
// 3. From the token payload, we use user_id to find form DB

func RequireAuth(appCtx component.AppContext) func(c *gin.Context) {

	tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}

		db := appCtx.GetMainDbConnection()
		store := storageuser.NewSQLStore(db)

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}
		user, err := store.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})
		if err != nil {
			panic(err)
		}

		if user.Active == 0 {
			panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
		}

		c.Set(common.CurrentUser, user)
		c.Next()
	}

}
