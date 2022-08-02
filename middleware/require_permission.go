package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"task1/common"
	"task1/component"
)

func RequirePermission(appCtx component.AppContext, permission string) func(c *gin.Context) {
	return func(c *gin.Context) {
		// get current user
		user := c.MustGet(common.CurrentUser).(common.Requester)

		id := user.GetUserId()

		// get db from appCtx
		db := appCtx.GetMainDbConnection()

		// each id_user => only role_code
		role_code := -1
		db = db.Table("user_role").Where("user_id", id).Find(&role_code)

		// relationship role and permission => many - many
		var permissions []int
		db = db.Table("role_permission").Where("role_code", role_code).Find(&permissions)
		var tmp string

		if err := db.Table("permission").Where("name_permission", permission).First(&tmp).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				panic(common.ErrNoPermission(err))
			}
			panic(err)
			return
		}
		c.Next()
	}
}
