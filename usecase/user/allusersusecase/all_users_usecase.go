package allusersusecase

import (
	"loans/domain"
	"loans/repository/userrepository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AllUsersUsecase struct {
	userRepo *userrepository.UserRepository
}

func NewAllUsersUsecase(userRepo *userrepository.UserRepository) *AllUsersUsecase {
	return &AllUsersUsecase{
		userRepo: userRepo,
	}
}

func (a *AllUsersUsecase) GetUsers(page, limit int, userID string) ([]*domain.User, int, error) {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, 0, domain.ErrInvalidUserID
	}

	user, err := a.userRepo.GetUserByID(objectID)
	if err != nil {
		return nil, 0, domain.ErrUserNotFoundByID
	}

	if user.Role == "user" {
		return nil, 0, domain.ErrOnlyAdminCanViewAllUsers
	}

	totalUsers, err := a.userRepo.GetUserCount()
	if err != nil {
		return nil, 0, err
	}

	pageCount := (totalUsers-1)/limit + 1
	if page > pageCount {
		return nil, 0, domain.ErrPageNotFound
	}

	users, err := a.userRepo.GetUsers(page, limit)
	if err != nil {
		return nil, 0, err
	}

	return users, pageCount, nil
}
