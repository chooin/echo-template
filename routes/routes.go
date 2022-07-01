package routes

import (
	"app/controllers/auth"
	"app/controllers/health"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	v1 := e.Group("/v1")
	{
		// auth
		v1Auth := v1.Group("/auth")
		{
			v1Auth.POST("/login", auth.Login)
		}

		v1.GET("/health", health.Get)
	}
}
