package product

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/config"
)

var productData = Products{
	Product{
		ProductName: "Mi",
		Price:       763,
		Description: "Mi",
		CategoryID:  1,
		StoreID:     1,
	},
	Product{
		ProductName: "Samsung",
		Price:       200,
		Description: "Samsung",
		CategoryID:  1,
		StoreID:     1,
	},
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	cfg, err := config.LoadConfig("config-local")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	//fill with dumy data
	ProductRepoInit(ProductRepoMock{
		products: productData,
	})
	ProductControllerDef(r, cfg)

	return r
}

