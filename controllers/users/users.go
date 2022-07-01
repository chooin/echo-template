package users

import (
	"app/requests"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Get(e echo.Context) error {
	var user requests.User
	if err := e.Bind(&user); err != nil {
		return e.String(404, "Not Found")
	}

	return e.String(http.StatusOK, "Halo")
}
