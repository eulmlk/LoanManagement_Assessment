package routes

import (
	"loans/bootstrap"
	"loans/delivery/controller/forgotpasswordcontroller"
	"loans/delivery/controller/loginusercontroller"
	"loans/delivery/controller/passwordresetcontroller"
	"loans/delivery/controller/registerusercontroller"
	"loans/delivery/controller/tokenrefreshcontroller"
	"loans/delivery/controller/verifyusercontroller"
	"loans/delivery/middlewares"
	"loans/repository/tokenrepository"
	"loans/repository/userrepository"
	"loans/usecase/forgotpasswordusecase"
	"loans/usecase/loginuserusecase"
	"loans/usecase/passwordresetusecase"
	"loans/usecase/registeruserusecase"
	"loans/usecase/tokenrefreshusecase"
	"loans/usecase/verifyuserusecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRoutes(client *mongo.Client) *gin.Engine {
	router := gin.Default()

	dbName, err := bootstrap.GetEnv("DATABASE_NAME")
	if err != nil {
		panic(err)
	}

	database := client.Database(dbName)

	// Dependency injection
	// Repositories
	userRepo := userrepository.NewUserRepository(database)
	tokenRepo := tokenrepository.NewTokenRepository(database)

	// Usecases
	registerUserUsecase := registeruserusecase.NewRegisterUserUseCase(userRepo)
	verifyUserUsecase := verifyuserusecase.NewVerifyUserUseCase(userRepo)
	loginUserUsecase := loginuserusecase.NewLoginUserUsecase(userRepo, tokenRepo)
	tokenRefreshUsecase := tokenrefreshusecase.NewTokenRefreshUsecase(userRepo, tokenRepo)
	forgotPasswordUsecase := forgotpasswordusecase.NewForgotPasswordUsecase(userRepo)
	passwordResetUsecase := passwordresetusecase.NewPasswordResetUsecase(userRepo)

	// Controllers
	registerUserController := registerusercontroller.NewRegisterUserController(registerUserUsecase)
	verifyUserController := verifyusercontroller.NewVerifyUserController(verifyUserUsecase)
	loginUserController := loginusercontroller.NewLoginUserController(loginUserUsecase)
	tokenRefreshController := tokenrefreshcontroller.NewTokenRefreshController(tokenRefreshUsecase)
	forgotPasswordController := forgotpasswordcontroller.NewForgotPasswordController(forgotPasswordUsecase)
	passwordResetController := passwordresetcontroller.NewPasswordResetController(passwordResetUsecase)

	// Public routes
	publicRoutes := router.Group("/")
	{
		publicRoutes.POST("/users/register", registerUserController.RegisterUser)
		publicRoutes.GET("/users/verify-email", verifyUserController.VerifyUser)
		publicRoutes.POST("/users/login", loginUserController.LoginUser)
		publicRoutes.POST("/users/password-reset", forgotPasswordController.ForgotPassword)
		publicRoutes.GET("/users/password-update", passwordResetController.ResetPassword)
	}

	// Refresh token route
	router.POST("/users/refresh-token", middlewares.AuthMiddleware("refresh"), tokenRefreshController.RefreshToken)

	return router
}
