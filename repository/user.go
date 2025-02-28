package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/siwarung/besw/config"
	"github.com/siwarung/besw/model"
	"github.com/siwarung/besw/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Membuat user baru
func CreateUser(user *model.User) (*mongo.InsertOneResult, error) {
	userCollection := config.DB.Collection("users")

	// Set waktu CreatedAt dan UpdatedAt dengan time.Time
	if user.UserID.IsZero() {
		user.UserID = primitive.NewObjectID()
	}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// Set default role jika kosong
	if user.Role == "" {
		user.Role = "admin"
	}

	// Simpan user ke database
	insertData := bson.M{
		"_id":        user.UserID,
		"username":   user.Username,
		"phone":      user.Phone,
		"password":   user.Password,
		"role":       user.Role,
		"created_at": primitive.NewDateTimeFromTime(user.CreatedAt),
		"updated_at": primitive.NewDateTimeFromTime(user.UpdatedAt),
	}

	result, err := userCollection.InsertOne(context.Background(), insertData)
	if err != nil {
		fmt.Println("Error Insert:", err)
		return nil, err
	}
	return result, nil
}

// Periksa apakah username sudah digunakan
func CheckUsername(username string) (bool, error) {
	userCollection := config.DB.Collection("users")
	count, err := userCollection.CountDocuments(context.Background(), bson.M{"username": username})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// Cari user berdasarkan username
func FindUserByUsername(username string) (*model.User, error) {
	userCollection := config.DB.Collection("users")

	var user model.User
	err := userCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user tidak ditemukan")
		}
		fmt.Println("Error saat mencari user:", err)
		return nil, err
	}
	return &user, nil
}

// Verifikasi user berdasarkan username dan password
func VerifyUser(username, password string) (*model.User, error) {
	user, err := FindUserByUsername(username)
	if err != nil {
		return nil, errors.New("username atau password salah")
	}

	// Verifikasi password menggunakan bcrypt
	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("username atau password salah")
	}

	return user, nil
}
