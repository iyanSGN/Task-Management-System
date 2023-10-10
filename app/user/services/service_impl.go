package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"main.go/app/user"
	"main.go/app/user/repository"
)

type serviceImpl struct {
	DB *gorm.DB
	Repository repository.Repository
}


func NewService(DB *gorm.DB, Repository repository.Repository) Service {
	return &serviceImpl{
		DB: DB,
		Repository: Repository,
	}
}

func(s *serviceImpl) GetAll(c echo.Context) ([]user.UserResponseDTO, error) {
	var userRes []user.UserResponseDTO

	result, err := s.Repository.GetAll(c,s.DB)
	if err != nil {
		return userRes, c.JSON(http.StatusInternalServerError, err)
	}

	for _, user := range result {
		userRes = append(userRes, user.ToResponse())
	}

	return userRes, nil
}

func(s *serviceImpl) GetByID(c echo.Context, id uint) (user.UserResponseDTO, error) {
	var userRes user.UserResponseDTO

	result, err := s.Repository.GetByID(c, s.DB,id)
	if err != nil {
		return userRes, c.JSON(http.StatusInternalServerError, err)
	}

	userRes = result.ToResponse()

	return userRes, nil
}