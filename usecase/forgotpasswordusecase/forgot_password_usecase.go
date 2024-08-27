package forgotpasswordusecase

import (
	"loans/config"
	"loans/domain"
	"loans/repository/userrepository"

	"go.mongodb.org/mongo-driver/mongo"
)

type ForgotPasswordUsecase struct {
	userRepo *userrepository.UserRepository
}

func NewForgotPasswordUsecase(userRepo *userrepository.UserRepository) *ForgotPasswordUsecase {
	return &ForgotPasswordUsecase{
		userRepo: userRepo,
	}
}

func (f *ForgotPasswordUsecase) ForgotPassword(email string, newPassword string) error {
	err := config.IsValidEmail(email)
	if err != nil {
		return err
	}

	err = config.IsStrongPassword(newPassword)
	if err != nil {
		return err
	}

	user, err := f.userRepo.GetUserByEmail(email)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.ErrUserNotFoundByEmail
		}

		return err
	}

	hashedPassword, err := config.HashPassword(newPassword)
	if err != nil {
		return err
	}

	resetClaims := &domain.ResetClaims{
		UserID:      user.ID.Hex(),
		NewPassword: hashedPassword,
	}

	resetToken, err := config.GenerateToken(resetClaims)
	if err != nil {
		return err
	}

	emailHeader := "Reset your password"
	emailBody := "Click the link below to reset your password:\n\n" + "http://localhost:8080/password-update?token=" + resetToken

	return config.SendEmail(email, emailHeader, emailBody, false)
}
