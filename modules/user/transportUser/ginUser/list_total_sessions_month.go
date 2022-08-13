package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task1/common"
	"task1/component"
	bizuser "task1/modules/user/bizUser"
	usermodel "task1/modules/user/modelUser"
	storageuser "task1/modules/user/storageUser"
)

func ListTotalSessionsMonth(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var filter usermodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}

		paging.Fullfill()
		store := storageuser.NewSQLStore(appCtx.GetMainDbConnection())
		biz := bizuser.NewUserListLoginBiz(store)

		response, err := biz.ListSessionsMonth(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(response))
	}
}
