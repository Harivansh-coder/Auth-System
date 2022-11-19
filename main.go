package main

import (
	"harry/auth_system/routers"

	"github.com/gofiber/fiber/v2"
)

// album represents data about a record album.

func main() {
	router := fiber.New()

	// routers
	router.Get("/", defaultRoute)

	routers.UserRouters(router)

	routers.LoginRouters(router)

	router.Listen(":8080")

}

func defaultRoute(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{"message": "This is the default route for Authentication API"})
}
