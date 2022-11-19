package controllers

import (
	"context"
	"harry/auth_system/models"
	"harry/auth_system/utils"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// Login is the controller for the login route

func Login(c *fiber.Ctx) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	var user models.LoginUser
	var foundUser models.User

	defer cancel()

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(models.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
	defer cancel()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.UserResponse{Status: http.StatusInternalServerError, Message: "user doesn't exit", Data: &fiber.Map{"data": err.Error()}})
	}

	msg := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if msg != nil {
		return c.Status(http.StatusUnauthorized).JSON(models.UserResponse{Status: http.StatusUnauthorized, Message: "credentials no valid", Data: &fiber.Map{"data": err.Error()}})
	}

	token, _ := utils.GenerateAllTokens(foundUser.Email, foundUser.Name, foundUser.Id.String())

	return c.JSON(&fiber.Map{"token": token, "tokenType": "bearer"})

}
