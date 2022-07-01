package users

import (
	"app/models"
	"app/pkg"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Get(e echo.Context) error {
	userId := e.Param("id")

	var user models.User
	pkg.DB().Model(&models.User{}).Where("id = ?", userId).First(&user)

	return e.JSON(http.StatusOK, user)
}
