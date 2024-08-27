package promoteusercontroller

import (
	"loans/domain"
	"loans/usecase/user/promoteuserusecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PromoteUserController struct {
	usecase *promoteuserusecase.PromoteUserUsecase
}

func NewPromoteUserController(usecase *promoteuserusecase.PromoteUserUsecase) *PromoteUserController {
	return &PromoteUserController{
		usecase: usecase,
	}
}

func (p *PromoteUserController) PromoteUser(ctx *gin.Context) {
	var request struct {
		UserID   string `json:"user_id"`
		Promoted bool   `json:"promoted"`
	}

	err := ctx.BindJSON(&request)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: "Invalid request",
			Error:   "Cannot parse request body",
		})
		return
	}

	if request.UserID == "" {
		ctx.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: "Invalid request",
			Error:   "User ID is required",
		})
		return
	}

	claims := ctx.MustGet("claims").(*domain.LoginClaims)
	promotionWord := "promote"
	if !request.Promoted {
		promotionWord = "demote"
	}

	err = p.usecase.PromoteUser(claims.UserID, request.UserID, request.Promoted)
	if err != nil {
		code := domain.GetStatus(err)

		if code == http.StatusInternalServerError {
			log.Println(err)

			ctx.JSON(code, domain.Response{
				Success: false,
				Message: "Internal server error",
				Error:   "Something went wrong",
			})
			return
		}

		ctx.JSON(code, domain.Response{
			Success: false,
			Message: "Failed to " + promotionWord + " user",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.Response{
		Success: true,
		Message: "User has been " + promotionWord + "d",
	})
}
