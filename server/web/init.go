package web

import (
	"github.com/gin-gonic/gin"
	"ziweiSgServer/db"
	"ziweiSgServer/server/web/controller"
	"ziweiSgServer/server/web/middleware"
)

func Init(router *gin.Engine) {
	//mysql init
	db.TestDB()

	initRouter(router)
}

func initRouter(router *gin.Engine) {
	router.Use(middleware.Cors())
	router.GET("/account/register", controller.DefaultAccountController.Register)
}
