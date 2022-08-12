package ginrole

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task1/common"
	"task1/component"
	"task1/modules/role/bizRole"
	storagerole "task1/modules/role/storageRole"
)

func DeleteRoleByAdmin(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		role_code := c.Param("id")

		store := storagerole.NewSQLStore(appCtx.GetMainDbConnection())
		biz := bizrole.NewDeleteRoleBiz(store)
		if err := biz.DeleteRole(c.Request.Context(), map[string]interface{}{"role_code": role_code}); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
