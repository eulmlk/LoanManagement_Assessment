package passwordresetcontroller

import (
	"loans/domain"
	"loans/usecase/passwordresetusecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PasswordResetController struct {
	usecase *passwordresetusecase.PasswordResetUsecase
}

func NewPasswordResetController(usecase *passwordresetusecase.PasswordResetUsecase) *PasswordResetController {
	return &PasswordResetController{usecase: usecase}
}

func (p *PasswordResetController) ResetPassword(ctx *gin.Context) {
	tokenString := ctx.Query("token")

	if tokenString == "" {
		ctx.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: "Invalid request",
			Error:   "Token is required",
		})
		return
	}

	err := p.usecase.ResetPassword(tokenString)
	if err != nil {
		code := domain.GetStatus(err)

		if code == http.StatusInternalServerError {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, domain.Response{
				Success: false,
				Message: "Internal server error",
				Error:   "Something went wrong while resetting password",
			})
			return
		}

		ctx.JSON(code, domain.Response{
			Success: false,
			Message: "Failed to reset password",
			Error:   err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, domain.Response{
		Success: true,
		Message: "Password reset successful",
	})
}
