package health

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Result struct {
	Status string `json:"status"`
}

func Get(e echo.Context) error {
	return e.JSON(http.StatusOK, &Result{
		Status: "UP",
	})
}
