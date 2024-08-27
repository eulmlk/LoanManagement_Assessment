package addrootusecase

import (
	"loans/bootstrap"
	"loans/config"
	"loans/domain"
	"loans/repository/userrepository"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type AddRootUsecase struct {
	userRepo *userrepository.UserRepository
}

func NewAddRootUsecase(userRepo *userrepository.UserRepository) *AddRootUsecase {
	return &AddRootUsecase{
		userRepo: userRepo,
	}
}

func (a *AddRootUsecase) AddRoot() error {
	rootUsername, err := bootstrap.GetEnv("ROOT_USERNAME")
	if err != nil {
		return err
	}

	err = config.IsValidUsername(rootUsername)
	if err != nil {
		return err
	}

	_, err = a.userRepo.GetUserByUsername(rootUsername)
	if err == nil {
		return nil
	}

	if err != mongo.ErrNoDocuments {
		return err
	}

	rootPassword, err := bootstrap.GetEnv("ROOT_PASSWORD")
	if err != nil {
		return err
	}

	err = config.IsStrongPassword(rootPassword)
	if err != nil {
		return err
	}

	rootEmail, err := bootstrap.GetEnv("ROOT_EMAIL")
	if err != nil {
		return err
	}

	err = config.IsValidEmail(rootEmail)
	if err != nil {
		return err
	}

	hashedPassword, err := config.HashPassword(rootPassword)
	if err != nil {
		return err
	}

	user := &domain.User{
		Username: rootUsername,
		Email:    rootEmail,
		Password: hashedPassword,
		JoinedAt: time.Now(),
		Role:     "root",
	}

	return a.userRepo.InsertUser(user)
}
