package applyloanusecase

import (
	"loans/domain"
	"loans/repository/loanrepository"
	"loans/repository/userrepository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ApplyLoanUseCase struct {
	userRepo *userrepository.UserRepository
	loanRepo *loanrepository.LoanRepository
}

func NewApplyLoanUseCase(userRepo *userrepository.UserRepository, loanRepo *loanrepository.LoanRepository) *ApplyLoanUseCase {
	return &ApplyLoanUseCase{
		userRepo: userRepo,
		loanRepo: loanRepo,
	}
}

func (a *ApplyLoanUseCase) ApplyLoan(userID string, amount int) (*domain.Loan, error) {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, domain.ErrInvalidUserID
	}

	_, err = a.userRepo.GetUserByID(objectID)
	if err == mongo.ErrNoDocuments {
		return nil, domain.ErrUserNotFoundByID
	}

	if err != nil {
		return nil, err
	}

	_, err = a.loanRepo.GetLoanByUserID(objectID)
	if err == nil {
		return nil, domain.ErrAlreadyHasLoan
	}

	if err != mongo.ErrNoDocuments {
		return nil, err
	}

	loan := &domain.Loan{
		UserID: objectID,
		Amount: amount,
		Status: "Pending",
	}

	err = a.loanRepo.InsertLoan(loan)
	if err != nil {
		return nil, err
	}

	return a.loanRepo.GetLoanByUserID(objectID)
}
