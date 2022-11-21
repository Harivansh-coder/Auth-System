package main

import (
	"harry/auth_system/middleware"
	"harry/auth_system/routers"

	"github.com/gofiber/fiber/v2"
)

// album represents data about a record album.

func main() {
	router := fiber.New()

	// router.Use(middleware.Authentication())

	// routers
	router.Get("/", defaultRoute)

	routers.LoginRouters(router)
	routers.SignUpRouters(router)

	router.Use(middleware.Authentication())

	routers.UserRouters(router)

	router.Listen(":8080")

}

func defaultRoute(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{"message": "This is the default route for Authentication API"})
}
