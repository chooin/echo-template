package users

import (
	"app/models"
	"app/models/response"
	"app/pkg/mysql"
	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	userId := c.Param("id")
	var user models.User
	mysql.Client.Model(&models.User{}).Where("id = ?", userId).First(&user)
	return response.Ok(c, user)
}
