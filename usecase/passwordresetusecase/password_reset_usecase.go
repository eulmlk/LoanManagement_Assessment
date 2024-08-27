package passwordresetusecase

import (
	"errors"
	"loans/config"
	"loans/domain"
	"loans/repository/userrepository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PasswordResetUsecase struct {
	userRepo *userrepository.UserRepository
}

func NewPasswordResetUsecase(userRepo *userrepository.UserRepository) *PasswordResetUsecase {
	return &PasswordResetUsecase{userRepo: userRepo}
}

func (p *PasswordResetUsecase) ResetPassword(tokenString string) error {
	resetClaims := &domain.ResetClaims{}
	err := config.ValidateToken(tokenString, resetClaims)
	if err != nil {
		return err
	}

	objectID, err := primitive.ObjectIDFromHex(resetClaims.UserID)
	if err != nil {
		return err
	}

	_, err = p.userRepo.GetUserByID(objectID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.ErrUserNotFoundByID
		}

		return err
	}

	update := bson.M{
		"password": resetClaims.NewPassword,
	}

	return p.userRepo.UpdateUser(objectID, update)
}
