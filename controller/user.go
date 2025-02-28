package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/siwarung/besw/model"
	"github.com/siwarung/besw/repository"
	"github.com/siwarung/besw/utils"
)

// handler untuk membuat user baru
func CreateUser(c *fiber.Ctx) error {
	var user model.User

	// Parsing body request ke struct User
	if err := c.BodyParser(&user); err != nil {
		fmt.Println("Error parsing body: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Input tidak valid",
			"error":   err.Error(),
		})
	}

	// Debugging: Print parsed user data
	fmt.Printf("Parsed User: %+v\n", user)

	// Validasi input user
	if err := utils.ValidateUserInput(user); err != nil {
		fmt.Println("Validasi gagal:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Cek apakah username sudah ada
	if exist, _ := repository.CheckUsername(user.Username); exist {
		fmt.Println("Username sudah digunakan:", user.Username)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Username sudah digunakan",
		})
	}

	// Hash password sebelum disimpan
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		fmt.Println("Hash password error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Terjadi kesalahan saat membuat user",
		})
	}
	user.Password = hashedPassword

	// Simpan user ke database
	_, err = repository.CreateUser(&user)
	if err != nil {
		fmt.Println("Error insert ke database:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Terjadi kesalahan saat membuat user",
		})
	}

	fmt.Println("User berhasil dibuat:", user.Username)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User berhasil dibuat",
		"user":   user,
	})
}

