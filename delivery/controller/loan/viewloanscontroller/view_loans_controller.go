package viewloanscontroller

import (
	"loans/domain"
	"loans/usecase/loan/viewloansusecase"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ViewLoansController struct {
	usecase *viewloansusecase.ViewLoansUsecase
}

func NewViewLoansController(usecase *viewloansusecase.ViewLoansUsecase) *ViewLoansController {
	return &ViewLoansController{
		usecase: usecase,
	}
}

func (v *ViewLoansController) ViewLoans(ctx *gin.Context) {
	pageString := ctx.DefaultQuery("page", "1")
	limitString := ctx.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageString)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: "Invalid page",
			Error:   "Page must be a number",
		})
		return
	}

	limit, err := strconv.Atoi(limitString)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: "Invalid limit",
			Error:   "Limit must be a number",
		})
		return
	}

	if page < 1 {
		ctx.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: "Invalid page",
			Error:   "Page must be greater than 0",
		})
		return
	}

	if limit < 1 || limit > 100 {
		ctx.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: "Invalid limit",
			Error:   "Limit must be between 1 and 100",
		})
		return
	}

	claims := ctx.MustGet("claims").(*domain.LoginClaims)

	users, pageCount, err := v.usecase.ViewLoans(page, limit, claims.UserID)
	if err != nil {
		code := domain.GetStatus(err)

		if code == http.StatusInternalServerError {
			log.Println(err)
			ctx.JSON(code, domain.Response{
				Success: false,
				Message: "Internal server error",
				Error:   "Cannot get users",
			})
			return
		}

		ctx.JSON(code, domain.Response{
			Success: false,
			Message: "Error getting users",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.Response{
		Success:   true,
		Message:   "Users found",
		Count:     len(users),
		PageCount: pageCount,
		Data:      users,
	})
}
