package response

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Data = map[string]interface{}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ExceptionResponse struct {
	Message string  `json:"message"`
	Errors  []Error `json:"errors"`
}

func Ok(c echo.Context, data any, code ...int) error {
	httpStatus := http.StatusOK
	if len(code) > 0 && code[0] >= 200 && code[0] < 300 {
		httpStatus = code[0]
	}
	return c.JSON(httpStatus, data)
}

// ClientException 客户端异常
func ClientException(c echo.Context, message string, errors []Error, code ...int) error {
	httpStatus := http.StatusBadRequest
	if len(code) > 0 && code[0] >= 400 && code[0] < 500 {
		httpStatus = code[0]
	}
	if errors == nil {
		errors = []Error{}
	}
	return c.JSON(httpStatus, &ExceptionResponse{
		Message: message,
		Errors:  errors,
	})
}

// ServerException 服务端异常
func ServerException(c echo.Context, message string) error {
	return c.JSON(http.StatusInternalServerError, &ExceptionResponse{
		Message: message,
		Errors:  []Error{},
	})
}
