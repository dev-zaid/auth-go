package api

import (
	"github.com/gofiber/fiber/v2"
	authController "github.com/dev-zaid/auth-go/api/auth/controller"
)

func handleRoot(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"message": "Choose a route to continue"})
}

func SetupApp(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/", handleRoot)
	api.Post("/login", authController.loginUser)
}
