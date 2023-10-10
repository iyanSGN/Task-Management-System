package repository

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"main.go/models"
)

type repositoryImpl struct {
}

func NewRepository() Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) GetAll(c echo.Context, DB *gorm.DB) ([]models.MUser, error) {
	var user []models.MUser

	if err := DB.Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
} 

func (r *repositoryImpl) GetByID(c echo.Context, DB *gorm.DB, id uint) (models.MUser, error) {
	var user models.MUser

	if err := DB.Where("id= ?", id).First(&user).Error;
	err != nil {
		return user, nil
	}

	return user, nil
}
