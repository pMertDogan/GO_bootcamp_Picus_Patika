package order


import (
	"encoding/json"
	"gorm.io/gorm"
	//https://github.com/go-gorm/datatypes
	"gorm.io/datatypes"

)

type Orders []Order

type Order struct {
	gorm.Model
	UserID int 
	 ProductID int
	ProductSnapshot datatypes.JSON //We can store snapshot of product as JSON
	Comment string
	ShippingAdress string
	BillingAddress string
	Quantity int
}

// fromJson Order
func UnmarshalOrders(data []byte) (Orders, error) {
	var r Orders
	err := json.Unmarshal(data, &r)
	return r, err
}


//Order toJson
func (r *Orders) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

//Order toJson
func (r *Order) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
