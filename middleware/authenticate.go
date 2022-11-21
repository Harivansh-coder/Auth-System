package middleware

// import (
// 	"fmt"
// 	"harry/auth_system/models"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/gofiber/fiber/v2"
// )

// // Authz validates token and authorizes users
// func Authentication() fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		clientToken := c.Request().Header.Peek("token")
// 		if clientToken == nil {
// 			c.JSON(http.StatusInternalServerError, models.LoginResponse{"error": fmt.Sprintf("No Authorization header provided")})
// 			return
// 		}

// 		claims, err := helper.ValidateToken(clientToken)
// 		if err != "" {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
// 			c.Abort()
// 			return
// 		}

// 		c.Set("email", claims.Email)
// 		c.Set("first_name", claims.First_name)
// 		c.Set("last_name", claims.Last_name)
// 		c.Set("uid", claims.Uid)

// 		c.Next()

// 	}
// }
