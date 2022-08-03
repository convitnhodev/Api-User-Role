package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"task1/common"
	"task1/component"
	bizuser "task1/modules/user/biz_user"

	storageuser "task1/modules/user/storage_user"
)

func GetUserByAdmin(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		user_id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(err)
		}

		store := storageuser.NewSQLStore(appCtx.GetMainDbConnection())
		biz := bizuser.NewGetUserBiz(store)

		user, err := biz.GetUser(c.Request.Context(), user_id)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessReponse(&user))
	}
}
