package controllers

import (
	"net/http"
	"sync"

	"github.com/Makhubedu/price-compare-server/models"
	"github.com/Makhubedu/price-compare-server/services"
	"github.com/gin-gonic/gin"
)

func SumSearchItems(c *gin.Context) {

	searchText, ok := c.GetQuery("search")

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})

		return
	}

	var wg sync.WaitGroup

	wg.Add(5)

	var pnp []models.ItemModel
	var pnpError error
	go func() {
		defer wg.Done()
		pnp, pnpError = services.PnPService(searchText)
	}()

	var game []models.ItemModel
	var gameError error
	go func() {
		defer wg.Done()
		game, gameError = services.GameService(searchText)
	}()

	var makro []models.ItemModel
	var makroError error
	go func() {
		defer wg.Done()
		makro, makroError = services.MakroService(searchText)
	}()

	var checkers []models.ItemModel
	var checkersError error
	go func() {
		defer wg.Done()
		checkers, checkersError = services.CheckersService(searchText)
	}()

	var shoprite []models.ItemModel
	var shopriteError error
	go func() {
		defer wg.Done()
		shoprite, shopriteError = services.ShopriteService(searchText)
	}()

	wg.Wait()

	if pnpError != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": pnpError.Error(),
		})
		return
	}

	if gameError != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": gameError.Error(),
		})
		return
	}

	if makroError != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": makroError.Error(),
		})
		return
	}

	if checkersError != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": checkersError.Error(),
		})
		return
	}

	if shopriteError != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": shopriteError.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"pnp":      pnp,
		"game":     game,
		"makro":    makro,
		"checkers": checkers,
		"shoprite": shoprite,
	})
}
