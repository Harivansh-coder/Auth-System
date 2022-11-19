package routers

import (
	"harry/auth_system/controllers"

	"github.com/gofiber/fiber/v2"
)

func LoginRouters(app *fiber.App) {
	app.Post("/login", controllers.Login)
}
