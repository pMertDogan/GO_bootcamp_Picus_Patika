package basket

import (
	"encoding/json"

	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/product"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/user"
	"gorm.io/gorm"
)

type Baskets []Basket

/*
We are storing user , prouct and totalQuantity in Basket
*/
type Basket struct {
	gorm.Model
	UserID        int             `json:"user_id"`
	User          user.User       `json:"user"`
	ProductID     int             `json:"product_id"`
	Product       product.Product `json:"product"`
	TotalQuantity int             `json:"totalQuantity" gorm:"min:1,max:10"`
}

// fromJson Order
func UnmarshalOrders(data []byte) (Baskets, error) {
	var r Baskets
	err := json.Unmarshal(data, &r)
	return r, err
}

//Order toJson
func (r *Baskets) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

//Order toJson
func (r *Basket) Marshal() ([]byte, error) {
	return json.Marshal(r)
}


//return baskets items basketIDs as int array
func (r *Baskets) GenerateProductIDTotalQuantityMap()  map[int]int{
	//create basket map key is product ID and value is totalQuantity
	basketMap := make(map[int]int)
	for _, basket := range *r {
		basketMap[basket.ProductID] = basket.TotalQuantity
	}
	return  basketMap
} 
