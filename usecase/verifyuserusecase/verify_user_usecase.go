package verifyuserusecase

import (
	"loans/config"
	"loans/domain"
	"loans/repository/userrepository"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type VerifyUserUseCase struct {
	userRepo *userrepository.UserRepository
}

func NewVerifyUserUseCase(userRepo *userrepository.UserRepository) *VerifyUserUseCase {
	return &VerifyUserUseCase{
		userRepo: userRepo,
	}
}

func (v *VerifyUserUseCase) VerifyUser(tokenString string) (*domain.User, error) {
	claims := &domain.RegisterClaims{}
	err := config.ValidateToken(tokenString, claims)
	if err != nil {
		log.Println(err)
		return nil, domain.ErrInvalidToken
	}

	_, err = v.userRepo.GetUserByEmail(claims.User.Email)
	if err == nil {
		return nil, domain.ErrEmailAlreadyExists
	}

	if err != mongo.ErrNoDocuments {
		return nil, err
	}

	_, err = v.userRepo.GetUserByUsername(claims.User.Email)
	if err == nil {
		return nil, domain.ErrUsernameAlreadyExists
	}

	if err != mongo.ErrNoDocuments {
		return nil, err
	}

	err = v.userRepo.InsertUser(&claims.User)
	if err != nil {
		return nil, err
	}

	addedUser, err := v.userRepo.GetUserByUsername(claims.User.Username)
	if err != nil {
		return nil, err
	}

	return addedUser, nil
}
