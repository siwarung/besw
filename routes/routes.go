package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/siwarung/besw/controller"
)

func URL(app *fiber.App) {
	// Default route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Tersambung ke server SiWarung")
	})

	// Route auth
	userRoute := app.Group("/api/auth")
	userRoute.Post("/register", controller.CreateUser)
	userRoute.Post("/login", controller.LoginUser)

	// Protected routes (Hanya bisa diakses dengan token JWT)
	// protected := app.Group("/api/auth", middleware.JWTMiddleware)
	// protected.Get("/profile", controller.ProfileUser)

	// Route 404 harus di paling bawah agar tidak mengganggu route lain
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).SendString("404 | Halaman tidak ditemukan")
	})
}
