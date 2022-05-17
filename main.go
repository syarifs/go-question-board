package main

import (
	"go-question-board/internal/core/service"
	"go-question-board/internal/framework/database"
	"go-question-board/internal/framework/repository"
	"go-question-board/internal/framework/routes"
	"go-question-board/internal/framework/transport/controller"
	"go-question-board/internal/framework/transport/middleware"
	"go-question-board/internal/utils"

	_ "go-question-board/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title           Question Board
// @version         1.0
// @description     server API for Question Board Application.

// @securityDefinitions.apikey ApiKey
// @in header
// @name Authorization

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api
func main() {

	utils.LoadConfig()

	db := database.InitDatabase(utils.DB_DRIVER)
	repo := repository.NewRepository(db)
	serv := service.NewService(repo)
	ctrl := controller.NewController(serv)

	e := echo.New()
	e.GET("/*", echoSwagger.WrapHandler)

	api := e.Group("/api")
	routes.NewRoutes(api, ctrl, middleware.JWT())
	
	middleware.Logging(e)

	e.Start(":" + utils.SERVER_PORT)
}
