package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"task1/component"
	"task1/middleware"
	"task1/modules/role/transport/gin_role"
	"task1/modules/user/transport/gin_user"
)

func main() {
	dsn := ("taskIbennefit:Thaothaothao2230@tcp(localhost:3306)/task1?charset=utf8mb4&parseTime=True&loc=Local")
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Println(db)
	db = db.Debug()

	if err := runService(db, "viethung"); err != nil {
		log.Fatalln(err)
	}
}

//func New() *TemplateRepo {
//	 db := database.InitDb()
//	 db.AutoMigrate(&models.Templatel{})
//	 return &TemplateRepo{Db: db}
//}

func runService(db *gorm.DB, secretKey string) error {
	r := gin.Default()
	// should use os variable
	appCtx := component.NewAppContext(db, secretKey)

	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/v1")

	user := v1.Group("/user")
	{
		user.POST("/new", ginuser.CreateUserByAdmin(appCtx))
		user.DELETE("/delete/:id", ginuser.DeleteUserByAdmin(appCtx))
		user.PATCH("/update/:id", ginuser.UpdateUserByAdmin(appCtx))
		user.GET("/get/:id", ginuser.GetUserByAdmin(appCtx))
		user.GET("/list/", ginuser.ListUserByAdmin(appCtx))
	}

	role := v1.Group("/role")
	{
		role.POST("/new", ginrole.CreateRoleByAdmin(appCtx))
		role.DELETE("/delete/:id", ginrole.DeleteRoleByAdmin(appCtx))
		role.GET("/get/:id", ginrole.GetRoleByAdmin(appCtx))

	}

	return r.Run()

}
