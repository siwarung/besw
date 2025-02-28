package repository

import (
	"context"
	"time"

	"github.com/siwarung/besw/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		collection: db.Collection("users"),
	}
}

// Membuat user baru
func (r *UserRepository) CreateUser(user *model.User) (*mongo.InsertOneResult, error) {
	// Set waktu untuk CreatedAt dan UpdatedAt
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// Insert data ke MongoDB
	result, err := r.collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Periksa apakah username sudah digunakan
func (r *UserRepository) CheckUsername(username string) (bool, error) {
	count, err := r.collection.CountDocuments(context.Background(), bson.M{"username": username})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
