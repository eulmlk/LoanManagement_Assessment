package profileusecase

import (
	"errors"
	"loans/domain"
	"loans/repository/userrepository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProfileUseCase struct {
	userRepo *userrepository.UserRepository
}

func NewProfileUseCase(userRepo *userrepository.UserRepository) *ProfileUseCase {
	return &ProfileUseCase{userRepo: userRepo}
}

func (p *ProfileUseCase) GetProfile(id string) (*domain.Profile, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, domain.ErrInvalidID
	}

	user, err := p.userRepo.GetUserByID(objectID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrUserNotFoundByID
		}

		return nil, err
	}

	profile := &domain.Profile{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Bio:       user.Bio,
		JoinedAt:  user.JoinedAt,
	}

	return profile, nil
}

func (p *ProfileUseCase) UpdateProfile(id string, profile domain.Profile) (*domain.Profile, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	user, err := p.userRepo.GetUserByID(objectID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrUserNotFoundByID
		}

		return nil, err
	}

	update := bson.M{}

	if profile.FirstName != "" {
		update["first_name"] = profile.FirstName
	}

	if profile.LastName != "" {
		update["last_name"] = profile.LastName
	}

	if profile.Bio != "" {
		update["bio"] = profile.Bio
	}

	err = p.userRepo.UpdateUser(objectID, update)
	if err != nil {
		return nil, err
	}

	newProfile := &domain.Profile{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		JoinedAt: user.JoinedAt,
	}

	if profile.FirstName != "" {
		newProfile.FirstName = profile.FirstName
	}

	if profile.LastName != "" {
		newProfile.LastName = profile.LastName
	}

	if profile.Bio != "" {
		newProfile.Bio = profile.Bio
	}

	return newProfile, nil
}
