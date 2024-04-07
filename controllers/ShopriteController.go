package controllers

import (
	"net/http"

	"github.com/Makhubedu/price-compare-server/services"
	"github.com/gin-gonic/gin"
)

// GetShopriteSearchedItems godoc
// @Summary Get shoprite searched items
// @Schemes
// @Description Get shoprite searched items
// @Param search query string true "Search text for filtering items"
// @Tags Shops
// @Accept json
// @Produce json
// @Success 200 {array} models.ItemModel
// @Router /shoprite [get]
func GetShopriteSearchedItems(c *gin.Context) {

	searchText, ok := c.GetQuery("search")

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})

		return
	}

	response, err := services.ShopriteService(searchText)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Searched items is not found",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, response)
}
