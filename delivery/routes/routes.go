package routes

import (
	"loans/bootstrap"
	"loans/delivery/controller/loan/applyloancontroller"
	"loans/delivery/controller/loan/approveloancontroller"
	"loans/delivery/controller/loan/loanstatuscontroller"
	"loans/delivery/controller/user/alluserscontroller"
	"loans/delivery/controller/user/deleteusercontroller"
	"loans/delivery/controller/user/forgotpasswordcontroller"
	"loans/delivery/controller/user/loginusercontroller"
	"loans/delivery/controller/user/passwordresetcontroller"
	"loans/delivery/controller/user/profilecontroller"
	"loans/delivery/controller/user/promoteusercontroller"
	"loans/delivery/controller/user/registerusercontroller"
	"loans/delivery/controller/user/tokenrefreshcontroller"
	"loans/delivery/controller/user/verifyusercontroller"
	"loans/delivery/middlewares"
	"loans/repository/loanrepository"
	"loans/repository/tokenrepository"
	"loans/repository/userrepository"
	"loans/usecase/loan/applyloanusecase"
	"loans/usecase/loan/approveloanusecase"
	"loans/usecase/loan/loanstatususecase"
	"loans/usecase/user/addrootusecase"
	"loans/usecase/user/allusersusecase"
	"loans/usecase/user/deleteuserusecase"
	"loans/usecase/user/forgotpasswordusecase"
	"loans/usecase/user/loginuserusecase"
	"loans/usecase/user/passwordresetusecase"
	"loans/usecase/user/profileusecase"
	"loans/usecase/user/promoteuserusecase"
	"loans/usecase/user/registeruserusecase"
	"loans/usecase/user/tokenrefreshusecase"
	"loans/usecase/user/verifyuserusecase"

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
	loanRepo := loanrepository.NewLoanRepository(database)

	// Usecases
	registerUserUsecase := registeruserusecase.NewRegisterUserUseCase(userRepo)
	verifyUserUsecase := verifyuserusecase.NewVerifyUserUseCase(userRepo)
	loginUserUsecase := loginuserusecase.NewLoginUserUsecase(userRepo, tokenRepo)
	tokenRefreshUsecase := tokenrefreshusecase.NewTokenRefreshUsecase(userRepo, tokenRepo)
	forgotPasswordUsecase := forgotpasswordusecase.NewForgotPasswordUsecase(userRepo)
	passwordResetUsecase := passwordresetusecase.NewPasswordResetUsecase(userRepo)
	profileUsecase := profileusecase.NewProfileUseCase(userRepo)
	applyLoanUsecase := applyloanusecase.NewApplyLoanUseCase(userRepo, loanRepo)
	loanStatusUsecase := loanstatususecase.NewLoanStatusUseCase(userRepo, loanRepo)
	addRootUsecase := addrootusecase.NewAddRootUsecase(userRepo)
	promoteUserUsecase := promoteuserusecase.NewPromoteUserUsecase(userRepo)
	allUsersUsecase := allusersusecase.NewAllUsersUsecase(userRepo)
	deleteUserUsecase := deleteuserusecase.NewDeleteUserUsecase(userRepo)
	approveLoanUsecase := approveloanusecase.NewApproveLoanUsecase(userRepo, loanRepo)

	// Add root user
	err = addRootUsecase.AddRoot()
	if err != nil {
		panic(err)
	}

	// Controllers
	registerUserController := registerusercontroller.NewRegisterUserController(registerUserUsecase)
	verifyUserController := verifyusercontroller.NewVerifyUserController(verifyUserUsecase)
	loginUserController := loginusercontroller.NewLoginUserController(loginUserUsecase)
	tokenRefreshController := tokenrefreshcontroller.NewTokenRefreshController(tokenRefreshUsecase)
	forgotPasswordController := forgotpasswordcontroller.NewForgotPasswordController(forgotPasswordUsecase)
	passwordResetController := passwordresetcontroller.NewPasswordResetController(passwordResetUsecase)
	profileController := profilecontroller.NewProfileController(profileUsecase)
	applyLoanController := applyloancontroller.NewApplyLoanController(applyLoanUsecase)
	loanStatusController := loanstatuscontroller.NewLoanStatusController(loanStatusUsecase)
	promoteUserController := promoteusercontroller.NewPromoteUserController(promoteUserUsecase)
	allUsersController := alluserscontroller.NewAllUsersController(allUsersUsecase)
	deleteUserController := deleteusercontroller.NewDeleteUserController(deleteUserUsecase)
	approveLoanController := approveloancontroller.NewApproveLoanController(approveLoanUsecase)

	// Public routes
	publicRoutes := router.Group("/")
	{
		publicRoutes.POST("/users/register", registerUserController.RegisterUser)
		publicRoutes.GET("/users/verify-email", verifyUserController.VerifyUser)
		publicRoutes.POST("/users/login", loginUserController.LoginUser)
		publicRoutes.POST("/users/password-reset", forgotPasswordController.ForgotPassword)
		publicRoutes.GET("/users/password-update", passwordResetController.ResetPassword)
		publicRoutes.GET("/users/profile/:id", profileController.GetProfile)
	}

	// Refresh token route
	router.POST("/users/refresh-token", middlewares.AuthMiddleware("refresh"), tokenRefreshController.RefreshToken)

	// Private routes
	privateRoutes := router.Group("/")
	privateRoutes.Use(middlewares.AuthMiddleware("access"))
	{
		privateRoutes.GET("/users/profile", profileController.GetMyProfile)
		privateRoutes.PATCH("/users/profile", profileController.UpdateProfile)
		privateRoutes.POST("/loans", applyLoanController.ApplyLoan)
		privateRoutes.GET("/loans", loanStatusController.GetLoanStatus)
		privateRoutes.POST("admin/users/promote", promoteUserController.PromoteUser)
		privateRoutes.GET("admin/users", allUsersController.GetUsers)
		privateRoutes.DELETE("admin/users/:id", deleteUserController.DeleteUser)
		privateRoutes.PATCH("admin/loans/:id/status", approveLoanController.ApproveLoan)
	}

	return router
}
