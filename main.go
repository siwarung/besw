package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/siwarung/besw/config"
	"github.com/siwarung/besw/routes"
)

func main() {
	// Koneksi ke DB
	config.ConnectDB()

	// Buat instance dari App
	app := fiber.New()

	// Middleware
	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	app.Use(logger.New(logger.Config{
		Format:     "${time} ${status} ${method} ${path}\n",
		TimeFormat: "15:04:05",
	}))

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Next()
	})	

	// Route
	routes.URL(app)

	// Gunakan port dari environment variable Heroku
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Starting server on port %s...\n", port)
	log.Fatal(app.Listen(":" + port))
}
