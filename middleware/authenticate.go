package middleware

import (
	"harry/auth_system/models"
	"harry/auth_system/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Authz validates token and authorizes users
func Authentication() fiber.Handler {
	return func(c *fiber.Ctx) error {
		clientToken := c.Get("token")

		if clientToken == "" {
			return c.Status(http.StatusBadRequest).JSON(models.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": clientToken}})
		}

		claims, err := utils.ValidateToken(clientToken)
		if err != "" {
			return c.Status(http.StatusUnauthorized).JSON(models.UserResponse{Status: http.StatusUnauthorized, Message: "invalid token", Data: &fiber.Map{"data": clientToken}})
		}

		c.Set("email", claims.Email)
		c.Set("name", claims.Name)
		c.Set("id", claims.Id)

		c.Next()

		return nil
	}
}
