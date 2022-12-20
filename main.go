package main

import (
	"log"
	"os"
	"time"

	api "github.com/dev-zaid/auth-go/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

var startTime time.Time

func rootFunction(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"message": "Welcome to the AUTH API"})
}

func healthCheck(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"message": "OK", "uptime": time.Since(startTime).String()})
}

func setupRoutes(app *fiber.App) {
	app.Get("/", rootFunction)
	app.Get("/health", healthCheck)
	api.SetupApp(app)
}

func init() {
	startTime = time.Now()
	err := godotenv.Load()
	if err != nil {
		log.Panicln(err)
	}
}

func main() {
	log.Println("Server Starting")
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	setupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf(`
	################################################
	🛡️  Server listening on port: %s 🛡️
	################################################
  `, port)

	app.Listen(":" + port)
}
