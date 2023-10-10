package services

import (
	"github.com/labstack/echo/v4"
	"main.go/app/user"
)

type Service interface {
	GetAll(c echo.Context) ([]user.UserResponseDTO, error)
	GetByID(c echo.Context, id uint) (user.UserResponseDTO, error)
}