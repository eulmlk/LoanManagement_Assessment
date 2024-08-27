package approveloanusecase

import (
	"loans/domain"
	"loans/repository/loanrepository"
	"loans/repository/userrepository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ApproveLoanUsecase struct {
	userRepo *userrepository.UserRepository
	loanRepo *loanrepository.LoanRepository
}

func NewApproveLoanUsecase(userRepo *userrepository.UserRepository, loanRepo *loanrepository.LoanRepository) *ApproveLoanUsecase {
	return &ApproveLoanUsecase{
		userRepo: userRepo,
		loanRepo: loanRepo,
	}
}

func (a *ApproveLoanUsecase) ApproveLoan(claimID, id string, approved bool) error {
	userID, err := primitive.ObjectIDFromHex(claimID)
	if err != nil {
		return domain.ErrInvalidUserID
	}

	loanID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.ErrInvalidLoanID
	}

	user, err := a.userRepo.GetUserByID(userID)
	if err != nil {
		return err
	}

	if user.Role == "user" {
		return domain.ErrOnlyAdminCanApprove
	}

	loan, err := a.loanRepo.GetLoanByID(loanID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.ErrLoanNotFoundByID
		}

		return err
	}

	if loan.Status != "Pending" {
		return domain.ErrLoanAlreadyApproved
	}

	update := bson.M{}
	if approved {
		update["status"] = "Approved"
	} else {
		update["status"] = "Rejected"
	}

	return a.loanRepo.UpdateLoan(loanID, update)
}
