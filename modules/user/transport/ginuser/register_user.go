package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task1/component"
	"task1/component/hasher"
	userbiz "task1/modules/user/biz"
	usermodel "task1/modules/user/model"
	storageuser "task1/modules/user/storage"
)

func Register(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := storageuser.NewSQLStore(appCtx.GetMainDbConnection())
		md5 := hasher.NewMD5Hash()
		biz := userbiz.NewRegisterBiz(store, md5)
		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{"data": "true"})
	}
}
