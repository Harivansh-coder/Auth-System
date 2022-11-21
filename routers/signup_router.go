package routers

import (
	"harry/auth_system/controllers"

	"github.com/gofiber/fiber/v2"
)

func SignUpRouters(app *fiber.App) {

	// create the user controller
	app.Post("/users/create", controllers.CreateUser)

}
