package product

type UpdateProductDTO struct {
	Sku         string `json:"sku" gorm:"unique" `
	ProductName string `json:"productName" `
	Description string `json:"description" `
	Color       string `json:"color"`
	Price       int    `json:"price" `
	StockCount  int    `json:"stockCount"`
	StoreID     int    `json:"storeId" `
	CategoryID 	int    `json:"categoryId" `
}
