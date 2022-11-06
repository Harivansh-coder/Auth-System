package controllers

import (
	"fmt"
	"net/http"

	models "harry/auth_system/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	fmt.Println("Hello")
	c.JSON(http.StatusOK, models.Users)
}

func GetUserByID(c *gin.Context) {
	fmt.Println("Hello")
	var id string = c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "User by ID", "id": id})
}

func CreateUser(c *gin.Context) {
	fmt.Println("Hello")
	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func UpdateUser(c *gin.Context) {
	fmt.Println("Hello")
	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

func DeleteUser(c *gin.Context) {
	fmt.Println("Hello")
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
