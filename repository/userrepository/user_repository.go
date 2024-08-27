package userrepository

import (
	"context"
	"loans/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	UserCollection *mongo.Collection
}

func NewUserRepository(database *mongo.Database) *UserRepository {
	return &UserRepository{
		UserCollection: database.Collection("users"),
	}
}

func (u *UserRepository) InsertUser(user *domain.User) error {
	_, err := u.UserCollection.InsertOne(context.Background(), user)
	return err
}

func (u *UserRepository) GetUserByID(id primitive.ObjectID) (*domain.User, error) {
	filter := bson.M{"_id": id}
	user := &domain.User{}

	err := u.UserCollection.FindOne(context.Background(), filter).Decode(user)
	return user, err
}

func (u *UserRepository) GetUserByUsername(username string) (*domain.User, error) {
	filter := bson.M{"username": username}
	user := &domain.User{}

	err := u.UserCollection.FindOne(context.Background(), filter).Decode(user)
	return user, err
}

func (u *UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	filter := bson.M{"email": email}
	user := &domain.User{}

	err := u.UserCollection.FindOne(context.Background(), filter).Decode(user)
	return user, err
}

func (u *UserRepository) UpdateUser(id primitive.ObjectID, userData bson.M) error {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": userData,
	}

	_, err := u.UserCollection.UpdateOne(context.Background(), filter, update)
	return err
}

func (u *UserRepository) DeleteUser(id string) error {
	filter := bson.M{"_id": id}
	_, err := u.UserCollection.DeleteOne(context.Background(), filter)
	return err
}
