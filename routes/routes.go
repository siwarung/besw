package routes

import (
	"github.com/siwarung/besw/controller"
	"github.com/siwarung/besw/repository"

	"github.com/gofiber/fiber/v2"
)

func URL(app *fiber.App) {
	// Default route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Tersambung ke server SiWarung")
	})

	// Route 404
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).SendString("404 | Halaman tidak ditemukan")
	})

	// Route auth
	userRoute := app.Group("/api/auth")
	userRepo := repository.NewUserRepository()
	userRoute.Post("/register", controller.NewUserController(*userRepo).CreateUser)

}
