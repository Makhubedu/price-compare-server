package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Makhubedu/price-compare-server/models"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

func PnPService(searchText string) (items []models.ItemModel, err error) {
	// Create a new context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Navigate to the URL
	url := os.Getenv("PNP") + searchText
	fmt.Println(url)
	if err := chromedp.Run(ctx, chromedp.Navigate(url)); err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Wait for the specified selector to be ready (indicating that the page has finished loading)
	if err := chromedp.Run(ctx, chromedp.WaitReady(".product-grid-item.list-mobile.ng-star-inserted")); err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Extract HTML content after JavaScript rendering
	var htmlContent string
	if err := chromedp.Run(ctx, chromedp.OuterHTML("html", &htmlContent)); err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Parse HTML content with goquery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Extract data using goquery
	doc.Find(".cx-page.ng-star-inserted > .cx-page-section .product-grid-item.list-mobile.ng-star-inserted").Each(func(i int, s *goquery.Selection) {
		// Extract product information for each item
		name, _ := s.Attr("data-cnstrc-item-name")
		price, _ := s.Attr("data-cnstrc-item-price")
		productImage, _ := s.Find("img").Attr("src")
		id, idExists := s.Attr("data-cnstrc-item-id")

		if idExists {
			// Create an ItemModel instance
			var item models.ItemModel

			// // Map fields from productData to ItemModel
			item.Id = id
			item.Name = name
			item.Price = price
			item.ProductImage = productImage

			// Append the item to the slice
			items = append(items, item)
		}

	})

	// Return items and nil error
	return items, nil
}
