package main

import (
	"app/config"
	"app/routes"
)

func main() {
	config.Config()
	_ = routes.Routes().Run("0.0.0.0:8080")
}
