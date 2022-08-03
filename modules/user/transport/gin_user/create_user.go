package gin_user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task1/common"
	"task1/component"
	hasher2 "task1/component/hasher"
	bizuser "task1/modules/user/biz_user"
	usermodel "task1/modules/user/model_user"
	storageuser "task1/modules/user/storage_user"
)

func CreateUserByAdmin(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := storageuser.NewSQLStore(appCtx.GetMainDbConnection())
		md5 := hasher2.NewMD5Hash()

		biz := bizuser.NewCreateUserBiz(store, md5)
		if err := biz.CreateNewUser(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessReponse(&data))
	}
}
