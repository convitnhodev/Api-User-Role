package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"task1/component"
	"task1/middleware"
	"task1/modules/user/transport/ginuser_admin"
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

func runService(db *gorm.DB, secretKey string) error {
	r := gin.Default()
	// should use os variable
	appCtx := component.NewAppContext(db, secretKey)

	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/v1")

	admin_user := v1.Group("/admin")
	{
		admin_user.POST("/register/new", ginuser_admin.CreateUserByAdmin(appCtx))
		admin_user.DELETE("/delete/:id", ginuser_admin.DeleteUserByAdmin(appCtx))
		admin_user.PATCH("/update/:id", ginuser_admin.UpdateUserByAdmin(appCtx))
	}

	//user := v1.Group("/users")
	//{
	//	user.POST("/register", ginuser.Register(appCtx))
	//	user.POST("/login", ginuser.Login(appCtx))
	//}

	return r.Run()

}
