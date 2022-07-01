package routes

import (
	"app/controllers/auth"
	"app/controllers/health"
	"app/controllers/users"
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

		// users
		v1Users := v1.Group("/users")
		{
			v1Users.GET("/:id", users.Get)
		}

		v1.GET("/health", health.Get)
	}
}
