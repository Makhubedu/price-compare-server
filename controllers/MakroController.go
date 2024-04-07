package controllers

import (
	"net/http"

	"github.com/Makhubedu/price-compare-server/services"
	"github.com/gin-gonic/gin"
)

// GetMakroSearchedItems godoc
// @Summary Get makro searched items
// @Schemes
// @Param search query string true "Search text for filtering items"
// @Description Get makro searched items
// @Tags Shops
// @Accept json
// @Produce json
// @Success 200 {array} models.ItemModel
// @Router /makro [get]
func GetMakroSearchedItems(c *gin.Context) {

	searchText, ok := c.GetQuery("search")

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})

		return
	}

	response, err := services.MakroService(searchText)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, response)
}
