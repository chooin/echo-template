package auth

import (
	"app/models/response"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	return response.Ok(c, "Halo")
}
