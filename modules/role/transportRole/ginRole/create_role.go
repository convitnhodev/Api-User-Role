package ginrole

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task1/common"
	"task1/component"
	bizrole "task1/modules/role/bizRole"
	modelrole "task1/modules/role/modelRole"
	storagerole "task1/modules/role/storageRole"
)

func CreateRoleByAdmin(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data modelrole.Role
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := storagerole.NewSQLStore(appCtx.GetMainDbConnection())
		biz := bizrole.NewCreateRoleBiz(store)
		if err := biz.CreateRole(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(&data))
	}
}
