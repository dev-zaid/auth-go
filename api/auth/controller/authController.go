package authController

import (
	"context"
	"os"

	authModel "github.com/dev-zaid/auth-go/api/auth/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(c *fiber.Ctx) error {
	var user authModel.User
	c.BodyParser(&user)

	adminCollection, err := database.GetCollection(os.Getenv("DB_NAME"), "Admin")
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var check authModel.User
	error := adminCollection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&check)
	if error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email does not exist",
		})
	}
	e := bcrypt.CompareHashAndPassword([]byte(check.Password), []byte(user.Password))
	if e != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Password is incorrect",
		})
	}
	token, err := authService.GenerateToken(user.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	refresh, err := authService.GenerateRefreshToken(user.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Login successful",
		"token":   token,
		"refresh": refresh,
		"email":   user.Email,
	})}