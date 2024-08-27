package deleteuserusecase

import (
	"loans/domain"
	"loans/repository/userrepository"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeleteUserUsecase struct {
	userRepo *userrepository.UserRepository
}

func NewDeleteUserUsecase(userRepo *userrepository.UserRepository) *DeleteUserUsecase {
	return &DeleteUserUsecase{
		userRepo: userRepo,
	}
}

func (d *DeleteUserUsecase) DeleteUser(claimID, userID string) error {
	myObjectID, err := primitive.ObjectIDFromHex(claimID)
	if err != nil {
		return domain.ErrInvalidUserID
	}

	user, err := d.userRepo.GetUserByID(myObjectID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.ErrUserNotFoundByID
		}

		return err
	}

	log.Println(user.Role)
	log.Println(claimID)
	log.Println(userID)
	if claimID != userID && user.Role == "user" {
		return domain.ErrOnlyAdminCanDelete
	}

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return domain.ErrInvalidUserIDDelete
	}

	targetUser, err := d.userRepo.GetUserByID(userObjectID)
	if err != nil {
		return domain.ErrUserNotFoundByID
	}

	if targetUser.Role != "user" && user.Role != "root" {
		return domain.ErrOnlyRootCanDelete
	}

	if targetUser.Role == "root" {
		return domain.ErrCantDeleteRoot
	}

	err = d.userRepo.DeleteUser(userObjectID)
	if err != nil {
		return err
	}

	return nil
}
