package main

import (
	"app/config"
	"app/middleware"
	"app/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Debug = config.AppDebug
	routes.Routes(e)
	e.Use(middleware.CORS())
	e.Logger.Fatal(e.Start(":8080"))
}
