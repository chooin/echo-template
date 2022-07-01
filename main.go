package main

import (
	"app/config"
	"app/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config.Config(e)
	routes.Routes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
