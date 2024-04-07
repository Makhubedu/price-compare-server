package routes

import (
	"fmt"
	"os"

	"github.com/Makhubedu/price-compare-server/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() {

	r := gin.Default()

	r.GET("/api/checkers", controllers.GetCheckersSearchedItems)
	r.GET("/api/shoprite", controllers.GetShopriteSearchedItems)
	r.GET("/api/woolworths", controllers.GetWoolworthsSearchedItems)
	r.GET("/api/game", controllers.GetGameSearchedItems)
	r.GET("/api/makro", controllers.GetMakroSearchedItems)
	r.GET("/api/pnp", controllers.GetPnPSearchedItems)
	r.GET("/api/sum", controllers.SumSearchItems)

	// gin.SetMode(gin.ReleaseMode)
	// r.ForwardedByClientIP = true
	// r.SetTrustedProxies([]string{"127.0.0.1", "192.168.1.2", "10.0.0.0/8"})
	// r.Use(gin.Logger())

	// Start the http server.
	r.Run(fmt.Sprint(":", os.Getenv("PORT")))
}
