package handler

import (
	"github.com/labstack/echo/v4"
	"main.go/app/user/controller"
	"main.go/app/user/repository"
	"main.go/app/user/services"
	"main.go/pkg/database"
)

type handlersUser struct {
	Controller controller.Controller
}

func NewHandlerUser()*handlersUser {
	mdr := repository.NewRepository()
	mds := services.NewService(database.GetDB(), mdr)

	return &handlersUser{
		Controller: controller.NewController(mds),
	}

}

func (hd *handlersUser) Route(g *echo.Group) {
	g.GET("", hd.Controller.GetAll)
	g.GET("/:id", hd.Controller.GetByID)
}