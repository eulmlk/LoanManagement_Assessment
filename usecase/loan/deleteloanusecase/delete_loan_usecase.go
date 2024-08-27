package deleteloanusecase

import (
	"loans/repository/loanrepository"
	"loans/repository/userrepository"
)

type DeleteLoanUsecase struct {
	loanRepo *loanrepository.LoanRepository
	userRepo *userrepository.UserRepository
}

func NewDeleteLoanUsecase(loanRepo *loanrepository.LoanRepository, userRepo *userrepository.UserRepository) *DeleteLoanUsecase {
	return &DeleteLoanUsecase{
		loanRepo: loanRepo,
		userRepo: userRepo,
	}
}
