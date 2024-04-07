package main

import (
	_ "github.com/Makhubedu/price-compare-server/docs"

	"github.com/Makhubedu/price-compare-server/initializers"
	"github.com/Makhubedu/price-compare-server/routes"
)

func init() {
	initializers.LoadEnvVariables()
}

// @title Price Compare API
// @version 1.0
// @description This is a price comparison API for various shops in South Africa

// @host localhost:8000
// @BasePath /api/v1
func main() {

	routes.RegisterRoutes()
}
