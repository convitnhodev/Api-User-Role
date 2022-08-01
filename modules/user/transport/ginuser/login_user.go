package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task1/common"
	"task1/component"
	"task1/component/hasher"
	"task1/component/tokenprovider/jwt"
	userbiz "task1/modules/user/biz"
	usermodel "task1/modules/user/model"
	storageuser "task1/modules/user/storage"
)

func Login(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			common.ErrInvalidRequest(err)
		}

		db := appCtx.GetMainDbConnection()

		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := storageuser.NewSQLStore(db)
		md5 := hasher.NewMD5Hash()
		biz := userbiz.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*30)

		account, err := biz.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			// error handling
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessReponse(account))

	}
}
