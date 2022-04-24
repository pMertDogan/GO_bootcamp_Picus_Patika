package basket

import "gorm.io/gorm"

type BasketSToResponseDTO []BasketToResponseDTO

type BasketToResponseDTO struct {
	gorm.Model
	UserID        int `json:"user_id"`
	TotalQuantity int `json:"totalQuantity" gorm:"min:1,max:10"`
	ProductID     int `json:"product_id"`
	// User          user.User       `json:"user"`
	Sku         string `json:"sku" gorm:"unique" `
	ProductName string `json:"productName" `
	Description string `json:"description" `
	Color       string `json:"color"`
	Price       int    `json:"price" `
	StockCount  int    `json:"stockCount"`
	//This one is our foreign key
	CategoryID int `json:"categoryId"`
	StoreID    int `json:"storeId" `
	//We cann add them if we use alaternative query to get STORE INFO
	// Name string
	// // Description string
	// Phone string
	// Email string
	// Address string
}


