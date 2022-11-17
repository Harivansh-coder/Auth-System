package controllers

import "github.com/gofiber/fiber/v2"

// Login is the controller for the login route

func Login(c *fiber.Ctx) error {
	// var user models.User

	// c.BindJSON(&user)
	// token, err := user.Login()
	// if err != nil {
	// 	c.JSON(401, gin.H{"error": "Invalid credentials"})
	// 	return
	// }

	token := "jnvienbrbnirtunuieifuvsiuebve4b8rtb89r98b4rbt48rtbejfnvjsdf"
	return c.JSON(&fiber.Map{"token": token, "tokenType": "bearer"})
}
