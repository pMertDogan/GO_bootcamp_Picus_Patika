package product

type ReqCreateDTO struct {
	Sku         string `json:"sku" binding:"required,alphanumunicode"`
	ProductName string `json:"productName" binding:"required,min=3,max=50"`
	Description string `json:"description" binding:"required,min=3,max=150"`
	Color       string `json:"color" binding:"required,min=3,max=50"`
	// https://stackoverflow.com/questions/9452897/how-to-decode-json-with-type-convert-from-string-to-float64
	Price      int `json:"price,string" binding:"required,numeric"`
	StockCount int `json:"stockCount,string" binding:"required,numeric"`
	//https://gorm.io/docs/belongs_to.html
	//This one is our foreign key
	CategoryID int `json:"categoryId,string" binding:"required,numeric"`
	//BTW we dont need hardcode reference:ID cause looks like its default
	// Category category.Category `json:"category"`                                  //`gorm:"foreignKey:CategoryID;references:ID"`
	StoreID  int               `json:"storeId,string" binding:"required,numeric"` //This one is our foreign key for store
	// Store    store.Store       `json:"store"`                                     //`gorm:"foreignKey:StoreID;references:ID"`
	//Owner + ID -> OwnerID , Example Store + ID -> StoreID
}
