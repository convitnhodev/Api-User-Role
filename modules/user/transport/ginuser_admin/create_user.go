package ginuser_admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task1/component"
	hasher2 "task1/component/hasher"
	"task1/modules/user/biz/adminc_role"
	usermodel "task1/modules/user/model"
	storageuser "task1/modules/user/storage"
)

func CreateUserByAdmin(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := storageuser.NewSQLStore(appCtx.GetMainDbConnection())
		md5 := hasher2.NewMD5Hash()

		biz := adminc_role.NewCreateUserBiz(store, md5)
		if err := biz.CreateNewUser(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{"mail": data.Email})
	}
}
