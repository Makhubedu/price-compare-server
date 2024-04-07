package services

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Makhubedu/price-compare-server/models"
	"github.com/gocolly/colly"
)

// CheckersService retrieves items from Checkers website based on searchText
func ShopriteService(searchText string) (items []models.ItemModel, err error) {
	// Instantiate a new collector
	c := colly.NewCollector()

	// OnHTML callback to extract data from each product frame
	c.OnHTML(".product-frame", func(e *colly.HTMLElement) {
		// Extract the value of the data-product-ga attribute
		productGA := e.Attr("data-product-ga")

		// Create a variable to hold the parsed JSON data
		var productData map[string]interface{}

		// Unmarshal the JSON data into a map[string]interface{}
		if err := json.Unmarshal([]byte(productGA), &productData); err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return
		}

		// Create an ItemModel instance
		var item models.ItemModel

		// Map fields from productData to ItemModel
		item.Id = fmt.Sprintf("%v", productData["id"])
		item.Name = fmt.Sprintf("%v", productData["name"])
		item.Price = fmt.Sprintf("%v", productData["price"])
		item.Brand = fmt.Sprintf("%v", productData["brand"])
		item.Category = fmt.Sprintf("%v", productData["category"])
		item.Variant = fmt.Sprintf("%v", productData["variant"])
		item.List = fmt.Sprintf("%v", productData["list"])
		item.UnitSalePrice = fmt.Sprintf("%v", productData["unit_sale_price"])
		item.Stock = fmt.Sprintf("%v", productData["stock"])
		item.ProductImage = fmt.Sprintf("%v", productData["product_image_url"])

		// Append the item to the slice
		items = append(items, item)
	})

	// Get Checkers URL from environment variable
	url := os.Getenv("SHOPRITE")

	// Visit the constructed URL
	if err := c.Visit(url + searchText); err != nil {
		return nil, err
	}

	// Return items and nil error
	return items, nil
}
