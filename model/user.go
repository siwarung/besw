package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	UserID    primitive.ObjectID `json:"user_id,omitempty" bson:"_id,omitempty"`
	Username  string             `json:"username" bson:"username"`
	Phone     string             `json:"phone" bson:"phone"`
	Password  string             `json:"password,omitempty" bson:"password"`
	Role      string             `json:"role" bson:"role"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at"`
}
