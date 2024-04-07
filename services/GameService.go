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

func GameService(searchText string) (items []models.ItemModel, err error) {
	// Create a new context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Navigate to the URL
	url := fmt.Sprintf("%s/l/search/?q=%s&t=%s", os.Getenv("GAME"), searchText, searchText)
	if err := chromedp.Run(ctx, chromedp.Navigate(url)); err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Wait for the specified selector to be ready (indicating that the page has finished loading)
	if err := chromedp.Run(ctx, chromedp.WaitReady("[data-testid=plp_flat_list]")); err != nil {
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
	doc.Find("[data-testid=plp_flat_list] .css-1dbjc4n.r-13awgt0.r-edyy15").Each(func(i int, s *goquery.Selection) {
		// Extract product information for each item
		name := s.Find(".css-901oao.css-cens5h").Text()
		price := s.Find(".css-901oao.r-g3a1by").Text()
		imageURL, _ := s.Find("[data-testid=network_image] img").Attr("src")

		// Create an item model for the current item
		item := models.ItemModel{
			Id:           fmt.Sprintf("game-%d", i),
			Name:         name,
			Price:        price,
			ProductImage: imageURL,
		}

		// Append the item to the items slice
		items = append(items, item)
	})

	// Return items and nil error
	return items, nil
}
