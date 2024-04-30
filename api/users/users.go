package users

import (
	"app/models"
	"app/models/response"
	"app/pkg"
	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	userId := c.Param("id")

	var user models.User
	pkg.MySQL().Model(&models.User{}).Where("id = ?", userId).First(&user)

	return response.Ok(c, user)
}
