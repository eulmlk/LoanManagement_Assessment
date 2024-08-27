package verifyusercontroller

import (
	"loans/domain"
	"loans/usecase/user/verifyuserusecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VerifyUserController struct {
	usecase *verifyuserusecase.VerifyUserUseCase
}

func NewVerifyUserController(usecase *verifyuserusecase.VerifyUserUseCase) *VerifyUserController {
	return &VerifyUserController{
		usecase: usecase,
	}
}

func (v *VerifyUserController) VerifyUser(ctx *gin.Context) {
	token := ctx.Query("token")

	if token == "" {
		ctx.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: "Could not verify user",
			Error:   "Token is required",
		})
		return
	}

	user, err := v.usecase.VerifyUser(token)
	if err != nil {
		code := domain.GetStatus(err)

		if code == http.StatusInternalServerError {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, domain.Response{
				Success: false,
				Message: "Internal server error",
			})
			return
		}

		ctx.JSON(code, domain.Response{
			Success: false,
			Message: "Could not verify user",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.Response{
		Success: true,
		Message: "User verified successfully",
		Data:    user,
	})
}
