package deleteusercontroller

import (
	"loans/domain"
	"loans/usecase/user/deleteuserusecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteUserController struct {
	usecase *deleteuserusecase.DeleteUserUsecase
}

func NewDeleteUserController(usecase *deleteuserusecase.DeleteUserUsecase) *DeleteUserController {
	return &DeleteUserController{
		usecase: usecase,
	}
}

func (d *DeleteUserController) DeleteUser(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*domain.LoginClaims)
	userID := ctx.Param("id")

	err := d.usecase.DeleteUser(claims.UserID, userID)
	if err != nil {
		code := domain.GetStatus(err)

		if code == http.StatusInternalServerError {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, domain.Response{
				Success: false,
				Message: "Internal server error",
				Error:   "Something went wrong",
			})
			return
		}

		ctx.JSON(code, domain.Response{
			Success: false,
			Message: "Delete user failed",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
