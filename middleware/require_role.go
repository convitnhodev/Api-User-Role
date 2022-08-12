package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"task1/common"
	"task1/component"
)

func RequireRoles(appCtx component.AppContext, permissions ...string) func(c *gin.Context) {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)

		for _, value := range permissions {
			if strings.Contains(u.GetPermissions(), value) {
				c.Next()
				return
			}
		}

		panic(common.ErrNoPermission(nil))
	}
}
