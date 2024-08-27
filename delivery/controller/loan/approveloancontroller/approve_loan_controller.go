package approveloancontroller

import (
	"loans/domain"
	"loans/usecase/loan/approveloanusecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApproveLoanController struct {
	usecase *approveloanusecase.ApproveLoanUsecase
}

func NewApproveLoanController(usecase *approveloanusecase.ApproveLoanUsecase) *ApproveLoanController {
	return &ApproveLoanController{
		usecase: usecase,
	}
}

func (a *ApproveLoanController) ApproveLoan(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*domain.LoginClaims)
	id := ctx.Param("id")

	var request struct {
		Approved bool `json:"approved"`
	}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: "Invalid request",
			Error:   "Invalid JSON body",
		})
		return
	}

	err = a.usecase.ApproveLoan(claims.UserID, id, request.Approved)
	if err != nil {
		code := domain.GetStatus(err)

		if code == http.StatusInternalServerError {
			log.Println(err)
			ctx.JSON(code, domain.Response{
				Success: false,
				Message: "Internal server error",
			})
			return
		}

		ctx.JSON(code, domain.Response{
			Success: false,
			Message: "Failed to approve loan",
			Error:   err.Error(),
		})
		return

	}

	ctx.JSON(http.StatusOK, domain.Response{
		Success: true,
		Message: "Loan approved",
	})
}
