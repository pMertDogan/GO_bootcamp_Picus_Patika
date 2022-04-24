package basket

type UpdateBasketDTO struct {
	TotalQuantity int `json:"totalQuantity,required"`
	BasketID 	int `json:"basketID,required"`
}