package registeruserusecase

import (
	"loans/bootstrap"
	"loans/config"
	"loans/domain"
	"loans/repository/userrepository"
)

type RegisterUserUseCase struct {
	userRepo *userrepository.UserRepository
}

func NewRegisterUserUseCase(userRepo *userrepository.UserRepository) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		userRepo: userRepo,
	}
}

func (r *RegisterUserUseCase) RegisterUser(user domain.User) error {
	err := config.IsValidUsername(user.Username)
	if err != nil {
		return err
	}

	err = config.IsValidEmail(user.Email)
	if err != nil {
		return err
	}

	err = config.IsStrongPassword(user.Password)
	if err != nil {
		return err
	}

	_, err = r.userRepo.GetUserByUsername(user.Username)
	if err == nil {
		return domain.ErrUsernameAlreadyExists
	}

	_, err = r.userRepo.GetUserByEmail(user.Email)
	if err == nil {
		return domain.ErrEmailAlreadyExists
	}

	user.Password, err = config.HashPassword(user.Password)
	if err != nil {
		return err
	}

	registerClaims := &domain.RegisterClaims{
		User: user,
	}

	tokenString, err := config.GenerateToken(registerClaims)
	if err != nil {
		return err
	}

	apiBaseURL, err := bootstrap.GetEnv("API_BASE_URL")
	if err != nil {
		return err
	}

	emailHeader := "Email Verification"
	emailBody := "Hi, " + user.Username + "! Click this link to verify your email: <a href='" + apiBaseURL + "/users/verify-email?token=" + tokenString + "'>Verify Email</a>"

	return config.SendEmail(user.Email, emailHeader, emailBody, true)
}
