package gin_user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"task1/common"
	"task1/component"
	bizuser "task1/modules/user/biz_user"

	storageuser "task1/modules/user/storage_user"
)

func DeleteUserByAdmin(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		user_id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(err)
		}

		store := storageuser.NewSQLStore(appCtx.GetMainDbConnection())
		biz := bizuser.NewDeleteUserBiz(store)
		if err := biz.DeleteUser(c.Request.Context(), user_id); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessReponse(true))
	}
}
