package utils

import (
	"errors"
	"regexp"

	"github.com/siwarung/besw/model"
)

// Regex untuk username (hanya huruf dan angka, 5-20 karakter)
var usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9]{5,20}$`)

// Regex untuk nomor telepon (hanya angka, 10-15 karakter, diawali dengan 62)
var phoneRegex = regexp.MustCompile(`^62[0-9]{8,15}$`)

// Validasi input user
func ValidateUserInput(user model.User) error {
	if !usernameRegex.MatchString(user.Username) {
		return errors.New("username harus terdiri dari huruf dan angka, 5-20 karakter")
	}

	if !phoneRegex.MatchString(user.Phone) {
		return errors.New("nomor telepon harus terdiri dari angka, 10-15 karakter, diawali dengan 62")
	}

	if len(user.Password) < 6 {
		return errors.New("password minimal 6 karakter")
	}

	if user.Role != "admin" && user.Role != "user" {
		return errors.New("role harus admin atau user")
	}

	return nil
}
