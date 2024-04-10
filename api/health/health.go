package health

import (
	"app/models/response"
	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	return response.Ok(c, response.Data{
		"status": "UP",
	})
}
