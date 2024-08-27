package profilecontroller

import (
	"loans/domain"
	"loans/usecase/user/profileusecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	usecase *profileusecase.ProfileUseCase
}

func NewProfileController(usecase *profileusecase.ProfileUseCase) *ProfileController {
	return &ProfileController{usecase: usecase}
}

func (p *ProfileController) GetProfile(ctx *gin.Context) {
	id := ctx.Param("id")

	profile, err := p.usecase.GetProfile(id)
	if err != nil {
		code := domain.GetStatus(err)

		if code == http.StatusInternalServerError {
			log.Println(err)
			ctx.JSON(code, domain.Response{
				Success: false,
				Message: "Internal server error",
				Error:   "Failed to get profile",
			})
			return
		}

		ctx.JSON(code, domain.Response{
			Success: false,
			Message: "Failed to get profile",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.Response{
		Success: true,
		Message: "Profile found",
		Data:    profile,
	})
}

func (p *ProfileController) GetMyProfile(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*domain.LoginClaims)
	profile, err := p.usecase.GetProfile(claims.UserID)
	if err != nil {
		code := domain.GetStatus(err)

		if code == http.StatusInternalServerError {
			log.Println(err)
			ctx.JSON(code, domain.Response{
				Success: false,
				Message: "Internal server error",
				Error:   "Failed to get profile",
			})
			return
		}

		ctx.JSON(code, domain.Response{
			Success: false,
			Message: "Failed to get profile",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.Response{
		Success: true,
		Message: "Profile found",
		Data:    profile,
	})
}

func (p *ProfileController) UpdateProfile(ctx *gin.Context) {
	var request struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Bio       string `json:"bio"`
	}

	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: "Failed to update profile",
			Error:   "Invalid request body",
		})
		return
	}

	claims := ctx.MustGet("claims").(*domain.LoginClaims)
	profile := domain.Profile{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Bio:       request.Bio,
	}

	updatedProfile, err := p.usecase.UpdateProfile(claims.UserID, profile)
	if err != nil {
		code := domain.GetStatus(err)

		if code == http.StatusInternalServerError {
			log.Println(err)
			ctx.JSON(code, domain.Response{
				Success: false,
				Message: "Internal server error",
				Error:   "Failed to update profile",
			})
			return
		}

		ctx.JSON(code, domain.Response{
			Success: false,
			Message: "Failed to update profile",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.Response{
		Success: true,
		Message: "Profile updated",
		Data:    updatedProfile,
	})
}
