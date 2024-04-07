package controllers

import (
	"net/http"

	"github.com/Makhubedu/price-compare-server/services"
	"github.com/gin-gonic/gin"
)

// GetCheckersSearchedItems godoc
// @Summary Get checkers searched items
// @Schemes
// @Param search query string true "Search text for filtering items"
// @Description Get	checkers searched items
// @Tags Shops
// @Accept json
// @Produce json
// @Success 200 {array} models.ItemModel
// @Router /checkers [get]
func GetCheckersSearchedItems(c *gin.Context) {

	searchText, ok := c.GetQuery("search")

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})

		return
	}

	response, err := services.CheckersService(searchText)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Searched items is not found",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, response)
}
