package routers

import (
	controller "harry/auth_system/controllers"

	"github.com/gin-gonic/gin"
)

func LoginRouters(router *gin.Engine) {
	router.POST("/login", controller.Login)
}
