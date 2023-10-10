package response

import "github.com/labstack/echo/v4"

type Response struct {
	Success bool        `json:"success" default:"true"`
	Message string      `json:"message" default:"success"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c echo.Context, code int, msg string, data interface{}) error {
	response := Response{
		Success: true,
		Message: msg,
		Data:    data,
	}

	return c.JSON(code, response)
}