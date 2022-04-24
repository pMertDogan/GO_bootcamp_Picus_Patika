package basket

type AddToBasketDTO struct {
	// UserID int `json:"userID" binding:"required,number"`
	ProductID     int `json:"productID" binding:"required,numeric"`
	TotalQuantity int `json:"totalQuantity" binding:"required,numeric"`
}
