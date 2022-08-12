package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"task1/common"
	"task1/component"
	hasher2 "task1/component/hasher"
	bizuser "task1/modules/user/bizUser"
	usermodel "task1/modules/user/modelUser"
	storageuser "task1/modules/user/storageUser"
)

func UpdateUserByAdmin(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(err)
		}

		var data usermodel.UserUpdate
		data.Id = user_id
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := storageuser.NewSQLStore(appCtx.GetMainDbConnection())
		md5 := hasher2.NewMD5Hash()

		biz := bizuser.NewUpdateUserBiz(store, md5)
		if err := biz.UpdateUser(c.Request.Context(), user_id, &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(&data))
	}
}
