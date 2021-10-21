package main

import (
	"github.com/gin-gonic/gin"
	"template/config"
	"template/controllers/auth"
)

func main() {
	config.Config()
	
	r := gin.Default()

	Auth := r.Group("/v1/auth")
	{
		Auth.GET("/login", auth.Login)
	}

	_ = r.Run("0.0.0.0:8080")
}
