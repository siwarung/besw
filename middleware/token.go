package middleware

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/siwarung/besw/model"
)

// Gunakan SECRET_KEY atau default "mysecretkey"
var jwtSecret = []byte(getSecretKey())

func getSecretKey() string {
	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		return "SECRET_KEY" // Default untuk debugging (ubah saat di production)
	}
	return secret
}

// GenerateToken untuk user
func GenerateToken(user *model.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token berlaku 24 jam

	claims := model.JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
		User: model.UserClaims{
			UserID:   user.UserID.Hex(),
			Username: user.Username,
			Phone:    user.Phone,
			Role:     user.Role,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// ValidateToken memverifikasi token JWT
func ValidateToken(tokenString string) (*model.JWTClaims, error) {
	// Parse token dengan metode HS256
	token, err := jwt.ParseWithClaims(tokenString, &model.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Pastikan token menggunakan metode HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, fiber.ErrUnauthorized
	}

	claims, ok := token.Claims.(*model.JWTClaims)
	if !ok || !token.Valid {
		return nil, fiber.ErrUnauthorized
	}

	return claims, nil
}

// Middleware untuk melindungi route dengan JWT
func JWTMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Token tidak ditemukan",
		})
	}

	// Hapus prefix "Bearer " jika ada
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	claims, err := ValidateToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Token tidak valid",
		})
	}

	// Simpan informasi user di context untuk digunakan di handler lain
	c.Locals("user", claims.User)

	return c.Next()
}
