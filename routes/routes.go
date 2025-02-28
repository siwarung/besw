package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/siwarung/besw/config"
	"github.com/siwarung/besw/controller"
	"github.com/siwarung/besw/repository"
)

func URL(app *fiber.App) {
	// Default route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Tersambung ke server SiWarung")
	})

	userRepo := repository.NewUserRepository(config.DB)
	userCtrl := controller.NewUserController(userRepo)

	// Route auth
	userRoute := app.Group("/api/auth")
	userRoute.Post("/register", userCtrl.CreateUser)

	// Route 404 harus di paling bawah agar tidak mengganggu route lain
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).SendString("404 | Halaman tidak ditemukan")
	})
}
