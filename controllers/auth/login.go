package auth

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Login(e echo.Context) error {
	return e.String(http.StatusOK, "Halo")
}
