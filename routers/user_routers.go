package routers

import (
	controller "harry/auth_system/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouters(router *gin.Engine) {
	// accessing the user controller
	router.GET("/users", controller.GetUsers)
	router.GET("/users/:id", controller.GetUserByID)

	// modify the user controller
	router.POST("/users/create", controller.CreateUser)
	router.PUT("/users/update/:id", controller.UpdateUser)
	router.DELETE("/users/delete/:id", controller.DeleteUser)

}
