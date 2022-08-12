package ginUserControl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task1/common"
	"task1/component"
)

func GetProfile(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		data := c.MustGet(common.CurrentUser).(common.Requester)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.GetUserId()))
	}
}
