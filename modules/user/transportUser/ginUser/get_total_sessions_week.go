package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task1/common"
	"task1/component"
	bizuser "task1/modules/user/bizUser"
	storageuser "task1/modules/user/storageUser"
)

func GetTotalSessionsWeek(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Param("email")

		store := storageuser.NewSQLStore(appCtx.GetMainDbConnection())
		biz := bizuser.NewUserLoginWeekBiz(store)
		total, err := biz.CountSessionsWeek(c.Request.Context(), email)
		if err != nil {
			panic(err)
		}
		type Response struct {
			Email string `json:"email"`
			Total int    `json:"total_sessions"`
		}
		response := Response{
			email,
			*total,
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(response))
	}
}
