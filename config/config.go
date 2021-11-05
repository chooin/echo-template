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
		err := godotenv.Load(".env")
		if err != nil {
			panic("failed to load dotenv")
		}
	} else {
		err := godotenv.Load(env)
		if err != nil {
			panic("failed to load dotenv")
		}
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
