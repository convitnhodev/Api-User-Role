package ginUserControl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task1/common"
	"task1/component"
	"task1/component/hasher"
	"task1/modules/userControl/bizUserControl"
	"task1/modules/userControl/modelUserControl"
	"task1/modules/userControl/storageUserControl"
)

func ChangePassword(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		data := c.MustGet(common.CurrentUser).(common.Requester)

		var source modelUserControl.BothPassword
		if err := c.ShouldBind(&source); err != nil {
			panic(err)
		}

		db := appCtx.GetMainDbConnection()
		hasher := hasher.NewMD5Hash()
		biz := bizUserControl.NewChangPasswordBiz(storageUserControl.NewSQLStore(db), hasher)
		if err := biz.ChangePassword(c, data, source.OldPassword, modelUserControl.PASSWORD(source.NewPassword)); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
