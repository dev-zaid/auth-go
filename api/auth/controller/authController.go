package authController

import "github.com/gofiber/fiber/v2"

func loginUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"message": "Login route"})
}