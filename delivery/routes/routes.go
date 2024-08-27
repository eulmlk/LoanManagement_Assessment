package routes

import (
	"loans/bootstrap"
	"loans/delivery/controller/loginusercontroller"
	"loans/delivery/controller/registerusercontroller"
	"loans/delivery/controller/verifyusercontroller"
	"loans/repository/tokenrepository"
	"loans/repository/userrepository"
	"loans/usecase/loginuserusecase.go"
	"loans/usecase/registeruserusecase"
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

	// Controllers
	registerUserController := registerusercontroller.NewRegisterUserController(registerUserUsecase)
	verifyUserController := verifyusercontroller.NewVerifyUserController(verifyUserUsecase)
	loginUserController := loginusercontroller.NewLoginUserController(loginUserUsecase)

	// Public routes
	publicRoutes := router.Group("/")
	{
		publicRoutes.POST("/users/register", registerUserController.RegisterUser)
		publicRoutes.GET("/users/verify-email", verifyUserController.VerifyUser)
		publicRoutes.POST("/users/login", loginUserController.LoginUser)
	}

	return router
}
