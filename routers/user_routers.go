package routers

import (
	"harry/auth_system/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRouters(app *fiber.App) {
	// accessing the user controller
	app.Get("/users", controllers.GetAllUsers)
	//router.GET("/users/:id", controller.GetUserByID)

	// modify the user controller
	app.Post("/users/create", controllers.CreateUser)
	//router.PUT("/users/update/:id", controller.UpdateUser)
	//router.DELETE("/users/delete/:id", controller.DeleteUser)

}
