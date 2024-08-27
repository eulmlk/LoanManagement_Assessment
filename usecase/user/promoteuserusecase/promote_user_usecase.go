package promoteuserusecase

import (
	"loans/domain"
	"loans/repository/userrepository"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PromoteUserUsecase struct {
	userRepo *userrepository.UserRepository
}

func NewPromoteUserUsecase(userRepo *userrepository.UserRepository) *PromoteUserUsecase {
	return &PromoteUserUsecase{
		userRepo: userRepo,
	}
}

func (p *PromoteUserUsecase) PromoteUser(claimID string, userID string, promoted bool) error {
	myObjectID, err := primitive.ObjectIDFromHex(claimID)
	if err != nil {
		log.Println(err)
		return domain.ErrInvalidID
	}

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println(err)
		return domain.ErrInvalidIDPromote
	}

	claim, err := p.userRepo.GetUserByID(myObjectID)
	if err != nil {
		log.Println(err)
		return domain.ErrUserNotFoundByID
	}

	if claim.Role != "root" {
		return domain.ErrOnlyRootCanPromote
	}

	user, err := p.userRepo.GetUserByID(userObjectID)
	if err != nil {
		log.Println(err)
		return domain.ErrUserNotFoundByID
	}

	if user.Role == "user" && !promoted {
		return domain.ErrAlreadyDemoted
	}

	if user.Role != "user" && promoted {
		return domain.ErrAlreadyPromoted
	}

	update := bson.M{}
	if promoted {
		update["role"] = "admin"
	} else {
		update["role"] = "user"
	}

	return p.userRepo.UpdateUser(userObjectID, update)
}
