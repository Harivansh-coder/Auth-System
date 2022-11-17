package routers

import (
	"github.com/gofiber/fiber/v2"
)

func LoginRouters(app *fiber.App) {
	app.Post("/login")
}
