package main

import (
	"app/config"
	"app/routes"
)

func main() {
	config.Config()
	routes.Run("0.0.0.0:8080")
}
