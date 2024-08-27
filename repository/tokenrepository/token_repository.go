package tokenrepository

import (
	"context"
	"loans/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TokenRepository struct {
	tokenCollection *mongo.Collection
}

func NewTokenRepository(database *mongo.Database) *TokenRepository {
	return &TokenRepository{
		tokenCollection: database.Collection("tokens"),
	}
}

func (t *TokenRepository) InsertToken(token *domain.Token) error {
	_, err := t.tokenCollection.InsertOne(context.Background(), token)
	return err
}

func (t *TokenRepository) GetTokenByUserAndDevice(userID primitive.ObjectID, deviceID string) (*domain.Token, error) {
	filter := bson.M{"user_id": userID, "device_id": deviceID}
	token := &domain.Token{}

	err := t.tokenCollection.FindOne(context.Background(), filter).Decode(token)
	return token, err
}

func (t *TokenRepository) DeleteTokenByID(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := t.tokenCollection.DeleteOne(context.Background(), filter)
	return err
}
