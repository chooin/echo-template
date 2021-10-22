package config

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

func Config()  {
	env := os.Getenv("MODE")
	if env == "" {
		godotenv.Load(".env")
	} else {
		godotenv.Load(env)
	}
	appDebug, err := strconv.ParseBool(os.Getenv("APP_DEBUG"))
	if err == nil {
		if appDebug {
			gin.SetMode(gin.DebugMode)
		} else {
			gin.SetMode(gin.ReleaseMode)
		}
	} else {
		gin.SetMode(gin.DebugMode)
	}
}
