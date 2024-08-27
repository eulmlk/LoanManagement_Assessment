package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username  string             `json:"username" bson:"username"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	FirstName string             `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName  string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Bio       string             `json:"bio,omitempty" bson:"bio,omitempty"`
	Role      string             `json:"role" bson:"role"`
	JoinedAt  time.Time          `json:"joined_at" bson:"joined_at"`
}

type Profile struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username  string             `json:"username" bson:"username"`
	Email     string             `json:"email" bson:"email"`
	FirstName string             `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName  string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Bio       string             `json:"bio,omitempty" bson:"bio,omitempty"`
	JoinedAt  time.Time          `json:"joined_at" bson:"joined_at"`
}
