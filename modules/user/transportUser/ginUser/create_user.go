package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task1/common"
	"task1/component"
	hasher2 "task1/component/hasher"
	bizuser "task1/modules/user/bizUser"
	usermodel "task1/modules/user/modelUser"
	storageuser "task1/modules/user/storageUser"
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
