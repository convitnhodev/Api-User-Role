package gin_role

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task1/common"
	"task1/component"
	"task1/modules/role/biz_role"
	"task1/modules/role/model_role"
	storagerole "task1/modules/role/storage_role"
)

func CreateRoleByAdmin(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model_role.Role
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := storagerole.NewSQLStore(appCtx.GetMainDbConnection())
		biz := biz_role.NewCreateRoleBiz(store)
		if err := biz.CreateRole(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessReponse(&data))
	}
}
