package loanstatususecase

import (
	"loans/domain"
	"loans/repository/loanrepository"
	"loans/repository/userrepository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LoanStatusUseCase struct {
	userRepo *userrepository.UserRepository
	loanRepo *loanrepository.LoanRepository
}

func NewLoanStatusUseCase(userRepo *userrepository.UserRepository, loanRepo *loanrepository.LoanRepository) *LoanStatusUseCase {
	return &LoanStatusUseCase{
		userRepo: userRepo,
		loanRepo: loanRepo,
	}
}

func (l *LoanStatusUseCase) GetLoanStatus(userID string) (*domain.Loan, error) {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, domain.ErrInvalidUserID
	}

	_, err = l.userRepo.GetUserByID(objectID)
	if err == mongo.ErrNoDocuments {
		return nil, domain.ErrUserNotFoundByID
	}

	if err != nil {
		return nil, err
	}

	loan, err := l.loanRepo.GetLoanByUserID(objectID)
	if err == mongo.ErrNoDocuments {
		return nil, domain.ErrLoanNotFoundByUserID
	}

	if err != nil {
		return nil, err
	}

	return loan, nil
}
