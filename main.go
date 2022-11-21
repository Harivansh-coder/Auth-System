package main

import (
	"harry/auth_system/middleware"
	"harry/auth_system/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// album represents data about a record album.

func main() {
	router := fiber.New()

	// Default config
	router.Use(cors.New())

	// routers
	router.Get("/", defaultRoute)

	// Login auth route
	routers.LoginRouters(router)

	// User creation route
	routers.SignUpRouters(router)

	// token verification route
	router.Use(middleware.Authentication())

	// User access route
	routers.UserRouters(router)

	router.Listen(":8080")

}

func defaultRoute(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{"message": "This is the default route for Authentication API"})
}
