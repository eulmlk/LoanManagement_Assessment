package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Token struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID      primitive.ObjectID `json:"user_id" bson:"user_id"`
	DeviceID    string             `json:"device_id" bson:"device_id"`
	TokenString string             `json:"token_string" bson:"token_string"`
}
