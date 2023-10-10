package routes

import (
	"github.com/labstack/echo/v4"
	"main.go/handler"
)

func SetupRoutes(g *echo.Group) {


	//ROUTES
	handler.NewHandlerUser().Route(g.Group("/user"))
}