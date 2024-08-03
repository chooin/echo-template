package response

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Data = map[string]interface{}

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Ok(c echo.Context, data any, code ...int) error {
	httpStatus := http.StatusOK

	if len(code) > 0 && code[0] >= 200 && code[0] < 300 {
		httpStatus = code[0]
	}

	return c.JSON(httpStatus, data)
}

func ClientException(c echo.Context, message string, data any, code ...int) error {
	httpStatus := http.StatusBadRequest

	if len(code) > 0 && code[0] >= 400 && code[0] < 500 {
		httpStatus = code[0]
	}

	return c.JSON(httpStatus, &Response{
		Message: message,
		Data:    data,
	})
}

func ServerException(c echo.Context, message string) error {
	return c.JSON(http.StatusInternalServerError, Data{
		"message": message,
	})
}
