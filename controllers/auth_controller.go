package controllers

import (
	"context"
	"harry/auth_system/models"
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
	var foundUser models.LoginUser

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
		return c.Status(http.StatusInternalServerError).JSON(models.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	msg := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	defer cancel()
	if msg != nil {
		return c.Status(http.StatusUnauthorized).JSON(models.UserResponse{Status: http.StatusInternalServerError, Message: "credentials no valid", Data: &fiber.Map{"data": err.Error()}})
	}

	token, refreshToken, _ := helper.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, foundUser.User_id)

	helper.UpdateAllTokens(token, refreshToken, foundUser.User_id)

	c.JSON(http.StatusOK, foundUser)

}
