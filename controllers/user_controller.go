package controllers

import (
	"context"
	"harry/auth_system/models"
	"net/http"
	"time"

	config "harry/auth_system/configs"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = config.GetCollection(config.DB, "users")
var validate = validator.New()

// func GetUsers(c *gin.Context) {
// 	fmt.Println("Hello")
// 	c.JSON(http.StatusOK)
// }

// func GetUserByID(c *gin.Context) {
// 	fmt.Println("Hello")
// 	var id string = c.Param("id")
// 	c.JSON(http.StatusOK, gin.H{"message": "User by ID", "id": id})
// }

func CreateUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(models.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newUser := models.User{
		Id:       primitive.NewObjectID(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(models.UserResponse{Status: http.StatusCreated, Message: "user created", Data: &fiber.Map{"data": result}})
}

// func UpdateUser(c *gin.Context) {
// 	fmt.Println("Hello")
// 	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
// }

// func DeleteUser(c *gin.Context) {
// 	fmt.Println("Hello")
// 	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
// }
