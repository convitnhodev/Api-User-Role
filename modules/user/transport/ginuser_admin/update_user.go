package ginuser_admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"task1/common"
	"task1/component"
	hasher2 "task1/component/hasher"
	"task1/modules/user/biz/adminc_role"
	usermodel "task1/modules/user/model"
	storageuser "task1/modules/user/storage"
)

func UpdateUserByAdmin(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(err)
		}

		var data usermodel.UserUpdate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storageuser.NewSQLStore(appCtx.GetMainDbConnection())
		md5 := hasher2.NewMD5Hash()

		biz := adminc_role.NewUpdateUserBiz(store, md5)
		if err := biz.UpdateUser(c.Request.Context(), user_id, &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessReponse(&data))
	}
}
