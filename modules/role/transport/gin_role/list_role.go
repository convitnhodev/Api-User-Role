package ginrole

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task1/common"
	"task1/component"
	"task1/modules/role/biz_role"
	"task1/modules/role/model_role"
	storagerole "task1/modules/role/storage_role"
)

func ListRoleByAdmin(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var filter model_role.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}

		// set default
		paging.Fullfill()

		store := storagerole.NewSQLStore(appCtx.GetMainDbConnection())

		biz := biz_role.NewListRoleBiz(store)
		data, err := biz.ListRole(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessReponse(data))
	}
}
