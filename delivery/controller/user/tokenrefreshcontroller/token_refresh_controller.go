package tokenrefreshcontroller

import (
	"loans/domain"
	"loans/usecase/user/tokenrefreshusecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenRefreshController struct {
	usecase *tokenrefreshusecase.TokenRefreshUsecase
}

func NewTokenRefreshController(usecase *tokenrefreshusecase.TokenRefreshUsecase) *TokenRefreshController {
	return &TokenRefreshController{usecase: usecase}
}

func (t *TokenRefreshController) RefreshToken(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*domain.LoginClaims)

	newToken, err := t.usecase.RefreshToken(claims)
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
			Message: "Failed to refresh token",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.Response{
		Success: true,
		Message: "Token refreshed",
		Data:    newToken,
	})
}
