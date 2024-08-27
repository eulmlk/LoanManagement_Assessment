package userrepository

import (
	"context"
	"loans/domain"

	"go.mongodb.org/mongo-driver/bson"
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

func (u *UserRepository) InsertUser(User *domain.User) error {
	_, err := u.UserCollection.InsertOne(context.Background(), User)
	return err
}

func (u *UserRepository) GetUserByID(ID string) (*domain.User, error) {
	filter := bson.M{"_id": ID}
	user := &domain.User{}

	err := u.UserCollection.FindOne(context.Background(), filter).Decode(user)
	return user, err
}

func (u *UserRepository) GetUserByUsername(Username string) (*domain.User, error) {
	filter := bson.M{"username": Username}
	user := &domain.User{}

	err := u.UserCollection.FindOne(context.Background(), filter).Decode(user)
	return user, err
}

func (u *UserRepository) GetUserByEmail(Email string) (*domain.User, error) {
	filter := bson.M{"email": Email}
	user := &domain.User{}

	err := u.UserCollection.FindOne(context.Background(), filter).Decode(user)
	return user, err
}

func (u *UserRepository) UpdateUser(User *domain.User) error {
	filter := bson.M{"_id": User.ID}
	update := bson.M{
		"$set": bson.M{
			"first_name": User.FirstName,
			"last_name":  User.LastName,
			"bio":        User.Bio,
		},
	}

	_, err := u.UserCollection.UpdateOne(context.Background(), filter, update)
	return err
}

func (u *UserRepository) DeleteUser(ID string) error {
	filter := bson.M{"_id": ID}
	_, err := u.UserCollection.DeleteOne(context.Background(), filter)
	return err
}
