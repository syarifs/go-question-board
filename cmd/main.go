package main

import (
	"go-question-board/internal/core/service"
	"go-question-board/internal/framework/database"
	"go-question-board/internal/framework/repository"
	"go-question-board/internal/framework/routes"
	"go-question-board/internal/framework/transport/controller"
	"go-question-board/internal/framework/transport/middleware"
	"go-question-board/internal/utils"

	"github.com/labstack/echo/v4"
)

func main() {

	utils.LoadConfig()

	db := database.InitDatabase(utils.DB_DRIVER)
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	
	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)
	authController := controller.NewAuthController(authService)

	e := echo.New()
	routes.NewUserRoutes(e, userController, middleware.JWT())
	routes.NewAuthRoutes(e, authController)

	middleware.Logging(e)

	e.Start(":" + utils.SERVER_PORT)
}
