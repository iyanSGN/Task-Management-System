package repository

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"main.go/models"
)

type Repository interface {
	GetAll(c echo.Context, DB *gorm.DB) ([]models.MUser, error)
	GetByID(c echo.Context, DB *gorm.DB, id uint) (models.MUser, error)
}