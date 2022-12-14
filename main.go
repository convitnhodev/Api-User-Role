package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"task1/component"
	"task1/middleware"
	ginrole "task1/modules/role/transportRole/ginRole"
	ginuser "task1/modules/user/transportUser/ginUser"
	"task1/modules/userControl/transportUserControl/ginUserControl"
)

func main() {

	dsn := os.Getenv("NameDB")
	//dsn := ("taskIbennefit:Thaothaothao2230@tcp(localhost:3306)/task1?charset=utf8mb4&parseTime=True&loc=Local")

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Println(db)
	db = db.Debug()

	if err := runService(db, "viethung", component.TimeJWT{2000000, 100000000000000}); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB, secretKey string, timeJWT component.TimeJWT) error {
	r := gin.Default()
	// should use os variable
	appCtx := component.NewAppContext(db, secretKey, timeJWT)

	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/v1")

	user := v1.Group("/user")
	{
		user.POST("/create", ginuser.CreateUserByAdmin(appCtx))
		user.DELETE("/delete/:id", ginuser.DeleteUserByAdmin(appCtx))
		user.PATCH("/update/:id", ginuser.UpdateUserByAdmin(appCtx))
		user.GET("/get/:id", ginuser.GetUserByAdmin(appCtx))
		user.GET("/list/", ginuser.ListUserByAdmin(appCtx))
		user.GET("/total-sessions/week/:email", ginuser.GetTotalSessionsWeek(appCtx))
		user.GET("/total-sessions/month/:email", ginuser.GetTotalSessionsMonth(appCtx))
		user.GET("/list-total-sessions/week", ginuser.ListTotalSessionsWeek(appCtx))
		user.GET("/list-total-sessions/month", ginuser.ListTotalSessionsWeek(appCtx))
	}

	role := v1.Group("/role")
	{
		role.POST("/create", ginrole.CreateRoleByAdmin(appCtx))
		role.DELETE("/delete/:id", ginrole.DeleteRoleByAdmin(appCtx))
		role.GET("/get/:id", ginrole.GetRoleByAdmin(appCtx))
		role.GET("/list/", ginrole.ListRoleByAdmin(appCtx))
	}

	userControl := v1.Group("/user-control")
	{
		userControl.POST("/login", ginUserControl.Login(appCtx))
		userControl.GET("/profile", middleware.RequireAuth(appCtx), middleware.RequireRoles(appCtx, "get profile"), ginUserControl.GetProfile(appCtx))
		userControl.POST("/change-password", middleware.RequireAuth(appCtx), ginUserControl.ChangePassword(appCtx))

	}
	return r.Run(":8080")
}
