package product

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"mime/multipart"
	"strconv"

	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/category"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/store"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Products []Product

type Product struct {
	gorm.Model
	Sku         string `json:"sku" gorm:"unique" `
	ProductName string `json:"productName" `
	Description string `json:"description" `
	Color       string `json:"color"`
	// https://stackoverflow.com/questions/9452897/how-to-decode-json-with-type-convert-from-string-to-float64
	Price      int `json:"price" `
	StockCount int `json:"stockCount"`
	//https://gorm.io/docs/belongs_to.html
	//This one is our foreign key
	CategoryID int `json:"categoryId"`
	//BTW we dont need hardcode reference:ID cause looks like its default
	Category category.Category `json:"category"`        //`gorm:"foreignKey:CategoryID;references:ID"`
	StoreID  int               `json:"storeId" ` //This one is our foreign key for store
	Store    store.Store       `json:"store"`           //`gorm:"foreignKey:StoreID;references:ID"`
	//Owner + ID -> OwnerID , Example Store + ID -> StoreID
}

// fromJson Product
func UnmarshalStore(data []byte) (Products, error) {
	var r Products
	err := json.Unmarshal(data, &r)
	return r, err
}

//Products toJson
func (r *Products) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

//Product toJson
func (r *Product) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

//Convert DTO to Product
func FromReqCreateDTO(reqProduct ReqCreateDTO) Product {
	return Product{
		ProductName: reqProduct.ProductName,
		Description: reqProduct.Description,
		Color:       reqProduct.Color,
		Price:       reqProduct.Price,
		CategoryID:  reqProduct.CategoryID,
		StockCount:  reqProduct.StockCount,
		StoreID:     reqProduct.StoreID,
		Sku:         reqProduct.Sku,
	}
}


//Convert csv data to Product
func ProductFromCSV(file *multipart.FileHeader) (Products, error) {
	var nProduct Product
	var nProductS Products

	csvFile, err := file.Open()
	defer csvFile.Close()

	if err != nil {
		zap.L().Error("error opening csv file", zap.Error(err))
		return nil, errors.New("error opening csv file" + err.Error())
	}
	//creating a *csv.reader
	reader := csv.NewReader(csvFile)
	//reading csv files via reader and method of reader

	csvData, err := reader.ReadAll()
	if err != nil {
		zap.L().Error("error reading csv file", zap.Error(err))
		return nil, errors.New("error reading csv file" + err.Error())
	}
	//in csv file, each comma represent a column, with this knowledge every row has 12 column and each column parse where they are belong.

	for i, each := range csvData {
		// fmt.Println(i, each)
		//Skip csv header row

		if i > 0 {
			nProduct.Sku = each[0]
			nProduct.ProductName = each[1]
			nProduct.Description = each[2]
			nProduct.Color = each[3]
			nProduct.Price, _ = strconv.Atoi(each[4])
			nProduct.StockCount, _ = strconv.Atoi(each[5])
			nProduct.CategoryID, _ = strconv.Atoi(each[6])
			nProduct.StoreID, _ = strconv.Atoi(each[7])
			nProductS = append(nProductS, nProduct)

		} else if i == 0 {
			//Check Headers
			if each[0] != "sku" || each[1] != "productName" || each[2] != "description" || each[3] != "color" ||
			 each[4] != "price" || each[5] != "stockCount" || each[6] != "categoryID" || each[7] != "storeID" {

				zap.L().Error("error reading csv file", zap.Error(err))
				return nil, errors.New("error reading csv file. Headers are not correct" + err.Error())
			}
		}

	}

	return nProductS, nil
}
