package loginuserusecase

import (
	"loans/config"
	"loans/domain"
	"loans/repository/tokenrepository"
	"loans/repository/userrepository"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type LoginUserUsecase struct {
	userRepo  *userrepository.UserRepository
	tokenRepo *tokenrepository.TokenRepository
}

func NewLoginUserUsecase(userRepo *userrepository.UserRepository, tokenRepo *tokenrepository.TokenRepository) *LoginUserUsecase {
	return &LoginUserUsecase{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
	}
}

func (l *LoginUserUsecase) LoginUser(usernameoremail, password, deviceID string) (string, string, error) {
	loginType := "username"
	err := config.IsValidUsername(usernameoremail)
	if err != nil {
		log.Println(err)
		loginType = "email"
		err = config.IsValidEmail(usernameoremail)
		if err != nil {
			log.Println(err)
			return "", "", domain.ErrInvalidUsernameEmailPassword
		}
	}

	var user *domain.User
	if loginType == "username" {
		user, err = l.userRepo.GetUserByUsername(usernameoremail)
		if err != nil {
			log.Println(err)
			return "", "", domain.ErrInvalidUsernameEmailPassword
		}
	} else {
		user, err = l.userRepo.GetUserByEmail(usernameoremail)
		if err != nil {
			log.Println(err)
			return "", "", domain.ErrInvalidUsernameEmailPassword
		}
	}

	err = config.ComparePassword(user.Password, password)
	if err != nil {
		log.Println(err)
		return "", "", domain.ErrInvalidUsernameEmailPassword
	}

	accessClaims := &domain.LoginClaims{
		UserID: user.ID.Hex(),
		Type:   "access",
	}

	accessString, err := config.GenerateToken(accessClaims)
	if err != nil {
		return "", "", err
	}

	existingToken, err := l.tokenRepo.GetTokenByUserAndDevice(user.ID, deviceID)
	if err == nil {
		err = l.tokenRepo.DeleteTokenByID(existingToken.ID)
		if err != nil {
			return "", "", err
		}
	} else if err != mongo.ErrNoDocuments {
		return "", "", err
	}

	refreshClaims := &domain.LoginClaims{
		UserID: user.ID.Hex(),
		Type:   "refresh",
	}

	refreshString, err := config.GenerateToken(refreshClaims)
	if err != nil {
		return "", "", err
	}

	token := &domain.Token{
		UserID:      user.ID,
		DeviceID:    deviceID,
		TokenString: refreshString,
	}

	err = l.tokenRepo.InsertToken(token)
	if err != nil {
		return "", "", err
	}

	return accessString, refreshString, nil
}
