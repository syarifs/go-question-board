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
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
func main() {

	utils.LoadConfig()

	db := database.InitDatabase(utils.DB_DRIVER)
	repo := repository.NewRepository(db)
	serv := service.NewService(repo)
	ctrl := controller.NewController(serv)

	e := echo.New()
	e.GET("/*", echoSwagger.WrapHandler)
	routes.NewRoutes(e, ctrl)
	
	middleware.Logging(e)

	e.Start(":" + utils.SERVER_PORT)
}
