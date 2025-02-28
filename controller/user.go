package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/siwarung/besw/model"
	"github.com/siwarung/besw/repository"
	"github.com/siwarung/besw/utils"
)

type UserController struct {
	userRepo *repository.UserRepository
}

func NewUserController(userRepo *repository.UserRepository) *UserController {
	return &UserController{userRepo: userRepo}
}

// handler untuk membuat user baru
func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	var user model.User

	// Parsing body request ke struct User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Input tidak valid",
		})
	}

	// Validasi input user
	if err := utils.ValidateUserInput(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Cek apakah username sudah ada
	if exist, _ := uc.userRepo.CheckUsername(user.Username); exist {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Username sudah digunakan",
		})
	}

	// Hash password sebelum disimpan
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Terjadi kesalahan saat membuat user",
		})
	}
	user.Password = hashedPassword

	// Simpan user ke database
	_, err = uc.userRepo.CreateUser(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Terjadi kesalahan saat membuat user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User berhasil dibuat",
		"user":    user,
	})

}
