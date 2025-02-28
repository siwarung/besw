package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/siwarung/besw/middleware"
	"github.com/siwarung/besw/repository"
)

func LoginUser(c *fiber.Ctx) error {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Format input tidak valid",
		})
	}

	// Verifikasi user
	user, err := repository.VerifyUser(req.Username, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Username atau password salah",
		})
	}

	// Buat token JWT
	token, err := middleware.GenerateToken(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal membuat token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login berhasil",
		"token":   token,
	})
}
