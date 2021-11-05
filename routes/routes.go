package routes

import (
	"app/controllers/auth"
	"app/controllers/health"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.New()

	v1 := r.Group("/v1")
	{
		// auth
		v1Auth := v1.Group("/auth")
		{
			v1Auth.POST("/login", auth.Login)
		}

		v1.GET("/health", health.Get)
	}

	return r
}
