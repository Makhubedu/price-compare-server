package main

import (
	"github.com/Makhubedu/price-compare-server/initializers"
	"github.com/Makhubedu/price-compare-server/routes"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {

	routes.RegisterRoutes()
}
