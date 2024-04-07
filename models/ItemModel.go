package models

type ItemModel struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Price         string `json:"price"`
	Brand         string `json:"brand"`
	Category      string `json:"category"`
	Variant       string `json:"variant"`
	List          string `json:"list"`
	UnitSalePrice string `json:"unit+sale_price"`
	Stock         string `json:"stock"`
	ProductImage  string `json:"product_image"`
}
