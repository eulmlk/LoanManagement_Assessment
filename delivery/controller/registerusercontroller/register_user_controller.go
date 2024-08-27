package registerusercontroller

import (
	"loans/domain"
	"loans/usecase/registeruserusecase"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RegisterUserController struct {
	usecase *registeruserusecase.RegisterUserUseCase
}

func NewRegisterUserController(usecase *registeruserusecase.RegisterUserUseCase) *RegisterUserController {
	return &RegisterUserController{usecase: usecase}
}

func (r *RegisterUserController) RegisterUser(ctx *gin.Context) {
	var request struct {
		Username  string `json:"username"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Bio       string `json:"bio"`
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

	if request.Username == "" {
		ctx.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: "Invalid request",
			Error:   "Username is required",
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

	if request.Password == "" {
		ctx.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: "Invalid request",
			Error:   "Password is required",
		})
		return
	}

	user := domain.User{
		Username:  request.Username,
		Email:     request.Email,
		Password:  request.Password,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Bio:       request.Bio,
		Role:      "user",
		JoinedAt:  time.Now(),
	}

	err = r.usecase.RegisterUser(user)
	if err != nil {
		code := domain.GetStatus(err)

		if code == http.StatusInternalServerError {
			log.Println(err)
			ctx.JSON(code, domain.Response{
				Success: false,
				Message: "Internal server error",
				Error:   "Failed to register user",
			})
			return
		}

		ctx.JSON(code, domain.Response{
			Success: false,
			Message: "Could not register user",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.Response{
		Success: true,
		Message: "Verification email has been sent",
	})
}
