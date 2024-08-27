package loginusercontroller

import (
	"log"
	"net/http"

	"loans/config"
	"loans/domain"
	"loans/usecase/loginuserusecase"

	"github.com/gin-gonic/gin"
)

type LoginUserController struct {
	usecase *loginuserusecase.LoginUserUsecase
}

func NewLoginUserController(usecase *loginuserusecase.LoginUserUsecase) *LoginUserController {
	return &LoginUserController{usecase: usecase}
}

func (controller *LoginUserController) LoginUser(ctx *gin.Context) {
	var request struct {
		UsernameOrEmail string `json:"username_or_email"`
		Password        string `json:"password"`
	}

	err := ctx.BindJSON(&request)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: "Invalid request",
			Error:   "Failed to parse request body",
		})
		return
	}

	if request.UsernameOrEmail == "" {
		ctx.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: "Invalid request",
			Error:   "Username or email is required",
		})
		return
	}

	if request.Password == "" {
		ctx.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: "Invalid request",
			Error:   "Password is required",
		})
		return
	}

	deviceID, err := config.GenerateDeviceID(ctx.GetHeader("User-Agent"), ctx.ClientIP())
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, domain.Response{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	accessToken, refreshToken, err := controller.usecase.LoginUser(request.UsernameOrEmail, request.Password, deviceID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, domain.Response{
		Success: true,
		Message: "Successfully logged in",
		Data: gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		},
	})
}
