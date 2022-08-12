package ginUserControl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task1/common"
	"task1/component"
	"task1/component/hasher"
	"task1/component/tokenprovider/jwt"
	"task1/modules/userControl/bizUserControl"
	"task1/modules/userControl/modelUserControl"
	"task1/modules/userControl/storageUserControl"
)

func Login(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData modelUserControl.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			common.ErrInvalidRequest(err)
		}

		db := appCtx.GetMainDbConnection()

		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := storageUserControl.NewSQLStore(db)
		md5 := hasher.NewMD5Hash()
		timeSet := bizUserControl.NewSetTime(appCtx.GetTimeJWT().TimeAccess, appCtx.GetTimeJWT().TimeRefresh)
		biz := bizUserControl.NewLoginBusiness(store, tokenProvider, md5, timeSet)

		account, err := biz.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			// error handling
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
