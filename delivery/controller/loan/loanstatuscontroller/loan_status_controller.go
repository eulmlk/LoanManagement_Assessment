package loanstatuscontroller

import (
	"loans/domain"
	"loans/usecase/loan/loanstatususecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoanStatusController struct {
	usecase *loanstatususecase.LoanStatusUseCase
}

func NewLoanStatusController(usecase *loanstatususecase.LoanStatusUseCase) *LoanStatusController {
	return &LoanStatusController{
		usecase: usecase,
	}
}

func (l *LoanStatusController) GetLoanStatus(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*domain.LoginClaims)

	loan, err := l.usecase.GetLoanStatus(claims.UserID)
	if err != nil {
		code := domain.GetStatus(err)

		if code == http.StatusInternalServerError {
			ctx.JSON(code, domain.Response{
				Success: false,
				Message: "Internal server error",
				Error:   "Failed to get loan status",
			})
			return
		}

		ctx.JSON(code, domain.Response{
			Success: false,
			Message: "Failed to get loan status",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.Response{
		Success: true,
		Message: "Successfully get loan status",
		Data:    loan,
	})
}
