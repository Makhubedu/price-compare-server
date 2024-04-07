package routes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Makhubedu/price-compare-server/controllers"
	docs "github.com/Makhubedu/price-compare-server/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RegisterRoutes registers routes for the API
func RegisterRoutes() {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	// API routes
	v1 := r.Group("/api/v1")
	v1.GET("/pnp", controllers.GetPnPSearchedItems)
	v1.GET("/makro", controllers.GetMakroSearchedItems)
	v1.GET("/game", controllers.GetGameSearchedItems)
	v1.GET("/shoprite", controllers.GetShopriteSearchedItems)
	v1.GET("/checkers", controllers.GetCheckersSearchedItems)
	v1.GET("/woolworths", controllers.GetWoolworthsSearchedItems)
	v1.GET("/sum", controllers.GetSumSearchItems)

	// Configure server settings
	gin.SetMode(gin.ReleaseMode)
	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{"127.0.0.1", "192.168.1.2", "10.0.0.0/8"})
	r.Use(gin.Logger())

	r.Use(CORSMiddleware())

	// Start the http server.
	r.Run(fmt.Sprint(":", os.Getenv("PORT")))
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
