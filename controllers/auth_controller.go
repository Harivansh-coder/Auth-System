package controllers

import (
	"github.com/gin-gonic/gin"
)

// Login is the controller for the login route

func Login(c *gin.Context) {
	// var user models.User

	// c.BindJSON(&user)
	// token, err := user.Login()
	// if err != nil {
	// 	c.JSON(401, gin.H{"error": "Invalid credentials"})
	// 	return
	// }

	token := "jnvienbrbnirtunuieifuvsiuebve4b8rtb89r98b4rbt48rtbejfnvjsdf"
	c.JSON(200, gin.H{"token": token, "tokenType": "bearer"})
}
