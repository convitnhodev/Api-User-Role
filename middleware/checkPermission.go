package middleware

import (
	"github.com/gin-gonic/gin"
	"task1/common"
	"task1/component"
)

// 1. Get token from header
// 2. Validate token and parse to payload
// 3. From the token payload, we use user_id to find form DB

func CheckPermission(appCtx component.AppContext, permission string) func(c *gin.Context) {

	return func(c *gin.Context) {
		data := c.MustGet(common.CurrentUser).(common.Requester)
		db := appCtx.GetMainDbConnection()
		id := data.GetUserId()
		var role int
		db = db.Table("user_role").Where("user_id", id).Find(&role)
		var per []int
		db = db.Table("role_permission").Where("permission_code", role).Find(&per)
		var name_permission []string
		db = db.Table("name_permission").Where("permission_code in (?)", per).Find(&name_permission)

		for _, value := range name_permission {
			if permission == value {
				c.Next()
			}
		}
		appErr := common.ErrInternal(err.(error))
		// AbortWithStatusJson is destructor function, We must stop
		c.AbortWithStatusJSON(appErr.StatusCode, appErr)
		// catch panic of gin
		panic(err)
		return
	}

}
