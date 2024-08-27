package applyloancontroller

import (
	"loans/domain"
	"loans/usecase/loan/applyloanusecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApplyLoanController struct {
	usecase *applyloanusecase.ApplyLoanUseCase
}

func NewApplyLoanController(usecase *applyloanusecase.ApplyLoanUseCase) *ApplyLoanController {
	return &ApplyLoanController{
		usecase: usecase,
	}
}

func (a *ApplyLoanController) ApplyLoan(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*domain.LoginClaims)

	var request struct {
		Amount int `json:"amount"`
	}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: "Invalid request body",
		})
		return
	}

	if request.Amount <= 0 {
		ctx.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: "Invalid amount",
		})
		return
	}

	loan, err := a.usecase.ApplyLoan(claims.UserID, request.Amount)
	if err != nil {
		code := domain.GetStatus(err)

		if code == http.StatusInternalServerError {
			log.Println(err)
			ctx.JSON(code, domain.Response{
				Success: false,
				Message: "Internal server error",
				Error:   "Failed to apply for loan",
			})
			return
		}

		ctx.JSON(code, domain.Response{
			Success: false,
			Message: "Cannot apply for loan",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, domain.Response{
		Success: true,
		Message: "Loan application successful",
		Data:    loan,
	})
}
