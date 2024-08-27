package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Loan struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
	Amount int                `json:"amount" bson:"amount"`
	Status string             `json:"status" bson:"status"`
}
