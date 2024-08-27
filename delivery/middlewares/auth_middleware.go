package middlewares

import (
	"loans/config"
	"loans/domain"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(tokenType string) gin.HandlerFunc {
	if tokenType != "access" && tokenType != "refresh" {
		panic("tokenType must be either 'access' or 'refresh'")
	}

	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, domain.Response{
				Success: false,
				Message: "Unauthorized",
				Error:   "Authorization header is required",
			})
			ctx.Abort()
			return
		}

		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 && strings.ToLower(authHeaderParts[0]) != "bearer" {
			ctx.JSON(http.StatusUnauthorized, domain.Response{
				Success: false,
				Message: "Unauthorized",
				Error:   "Invalid authorization header format",
			})
			ctx.Abort()
			return
		}

		var claims domain.Claims
		if tokenType == "access" {
			claims = &domain.LoginClaims{Type: "access"}
		} else {
			claims = &domain.LoginClaims{Type: "refresh"}
		}

		tokenString := authHeaderParts[1]
		err := config.ValidateToken(tokenString, claims)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, domain.Response{
				Success: false,
				Message: "Unauthorized",
				Error:   err.Error(),
			})
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)
		ctx.Next()
	}
}
