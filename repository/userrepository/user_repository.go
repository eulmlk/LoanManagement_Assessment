package userrepository

import (
	"context"
	"loans/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (u *UserRepository) DeleteUser(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := u.UserCollection.DeleteOne(context.Background(), filter)
	return err
}

func (u *UserRepository) GetUsers(page, limit int) ([]*domain.User, error) {
	skip := (page - 1) * limit

	var users []*domain.User

	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(limit))

	cursor, err := u.UserCollection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		user := &domain.User{}

		err := cursor.Decode(user)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	err = cursor.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepository) GetUserCount() (int, error) {
	count, err := u.UserCollection.CountDocuments(context.Background(), bson.M{})
	return int(count), err
}
