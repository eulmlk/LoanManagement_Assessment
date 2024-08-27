package tokenrefreshusecase

import (
	"loans/config"
	"loans/domain"
	"loans/repository/tokenrepository"
	"loans/repository/userrepository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TokenRefreshUsecase struct {
	userRepo  *userrepository.UserRepository
	tokenRepo *tokenrepository.TokenRepository
}

func NewTokenRefreshUsecase(userRepo *userrepository.UserRepository, tokenRepo *tokenrepository.TokenRepository) *TokenRefreshUsecase {
	return &TokenRefreshUsecase{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
	}
}

func (t *TokenRefreshUsecase) RefreshToken(claims *domain.LoginClaims) (string, error) {
	if claims.Type != "refresh" {
		return "", domain.ErrInvalidToken
	}

	objectID, err := primitive.ObjectIDFromHex(claims.UserID)
	if err != nil {
		return "", err
	}

	_, err = t.userRepo.GetUserByID(objectID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", domain.ErrUserNotFoundByID
		}

		return "", err
	}

	claims.Type = "access"
	return config.GenerateToken(claims)
}
