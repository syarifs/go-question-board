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
	repo := repository.NewRepository(db)
	serv := service.NewService(repo)
	ctrl := controller.NewController(serv)

	e := echo.New()
	routes.NewRoutes(e, ctrl)
	
	middleware.Logging(e)

	e.Start(":" + utils.SERVER_PORT)
}
