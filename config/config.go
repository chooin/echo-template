package config

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func Config()  {
	env := os.Getenv("MODE")
	if env == "" {
		godotenv.Load(".env.debug")
	} else {
		godotenv.Load(env)
	}
	gin.SetMode(os.Getenv("GIN_MODE"))
}
