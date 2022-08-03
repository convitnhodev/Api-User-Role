package ginuser_admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task1/common"
	"task1/component"
	"task1/modules/user/biz/adminc_role"
	usermodel "task1/modules/user/model"
	storageuser "task1/modules/user/storage"
)

func ListUserByAdmin(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var filter usermodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}

		// set default
		paging.Fullfill()

		store := storageuser.NewSQLStore(appCtx.GetMainDbConnection())

		biz := adminc_role.NewListUserBiz(store)
		data, err := biz.ListUser(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessReponse(data))
	}
}
