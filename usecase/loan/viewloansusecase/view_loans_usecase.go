package viewloansusecase

import (
	"loans/domain"
	"loans/repository/loanrepository"
	"loans/repository/userrepository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ViewLoansUsecase struct {
	userRepo *userrepository.UserRepository
	loanRepo *loanrepository.LoanRepository
}

func NewViewLoansUsecase(userRepo *userrepository.UserRepository, loanRepo *loanrepository.LoanRepository) *ViewLoansUsecase {
	return &ViewLoansUsecase{
		userRepo: userRepo,
		loanRepo: loanRepo,
	}
}

func (v *ViewLoansUsecase) ViewLoans(page, limit int, claimID string) ([]*domain.Loan, int, error) {
	objectID, err := primitive.ObjectIDFromHex(claimID)
	if err != nil {
		return nil, 0, err
	}

	user, err := v.userRepo.GetUserByID(objectID)
	if err != nil {
		return nil, 0, domain.ErrUserNotFoundByID
	}

	if user.Role == "user" {
		return nil, 0, domain.ErrOnlyAdminCanViewLoans
	}

	totalLoans, err := v.loanRepo.GetLoanCount()
	if err != nil {
		return nil, 0, err
	}

	pageCount := (totalLoans-1)/limit + 1
	if page > pageCount {
		return nil, 0, domain.ErrPageNotFound
	}

	loans, err := v.loanRepo.GetLoans(page, limit)
	if err != nil {
		return nil, 0, err
	}

	return loans, pageCount, nil
}
