package config

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"os"
	"strconv"
)

func Config(e *echo.Echo) {
	env := os.Getenv("MODE")
	if env == "" {
		if err := godotenv.Load(".env"); err != nil {
			panic("Failed to load dotenv")
		}
	} else {
		if err := godotenv.Load(env); err != nil {
			panic("Failed to load dotenv")
		}
	}
	appDebug, err := strconv.ParseBool(os.Getenv("APP_DEBUG"))
	if err == nil {
		if appDebug {
			e.Debug = true
		}
	} else {
		e.Debug = true
	}
	e.Debug = false
}
