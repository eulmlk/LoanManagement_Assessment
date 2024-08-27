package forgotpasswordcontroller

import (
	"loans/domain"
	"loans/usecase/forgotpasswordusecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ForgotPasswordController struct {
	usecase *forgotpasswordusecase.ForgotPasswordUsecase
}

func NewForgotPasswordController(usecase *forgotpasswordusecase.ForgotPasswordUsecase) *ForgotPasswordController {
	return &ForgotPasswordController{
		usecase: usecase,
	}
}

func (f *ForgotPasswordController) ForgotPassword(ctx *gin.Context) {
	var request struct {
		Email       string `json:"email"`
		NewPassword string `json:"new_password"`
	}

	err := ctx.BindJSON(&request)
	if err != nil {
		log.Println(err)

		ctx.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: "invalid request body",
			Error:   "Failed to parse request body",
		})
		return
	}

	if request.Email == "" {
		ctx.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: "Invalid request",
			Error:   "Email is required",
		})
		return
	}

	if request.NewPassword == "" {
		ctx.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: "Invalid request",
			Error:   "New password is required",
		})
		return
	}

	err = f.usecase.ForgotPassword(request.Email, request.NewPassword)
	if err != nil {
		status := domain.GetStatus(err)

		if status == http.StatusInternalServerError {
			log.Println(err)
			ctx.JSON(status, domain.Response{
				Success: false,
				Message: "Internal server error",
				Error:   "Failed to send reset password email",
			})
			return
		}

		ctx.JSON(status, domain.Response{
			Success: false,
			Message: "failed to send reset password email",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.Response{
		Success: true,
		Message: "reset password email sent",
	})
}
