package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"main.go/app/user/services"
	"main.go/pkg/response"
)

type controllerImpl struct {
	Service services.Service
}

func NewController(Service services.Service) Controller {
	return &controllerImpl{
		Service: Service,
	}
}

func (co *controllerImpl)GetAll(c echo.Context) error {
	result, err := co.Service.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	return response.SuccessResponse(c, http.StatusOK, "Success Get All Users", result)
}

func (co *controllerImpl)GetByID(c echo.Context) error {
	userID := c.Param("id")

	id, err := strconv.ParseUint(userID, 10, 64)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	result, err := co.Service.GetByID(c, uint(id))
	if err != nil {
		return err
	}

	return response.SuccessResponse(c, http.StatusOK, "success get user by id", result)
}