package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain"
	"go.uber.org/zap"
)

func CreateProduct(c *gin.Context) {
	//Json
	c.Header("Content-Type", "application/json")
	
	var reqProduct ReqCreateDTO
	var responseModel domain.ResponseModel

	if err := c.Bind(&reqProduct); err != nil {
		responseModel.ErrMsg = err.Error()
		responseModel.ErrDsc = "Error while binding request body"
		responseModel.ResponseCode = http.StatusBadRequest
		c.JSON(responseModel.ResponseCode, responseModel)
		return
	}
	zap.L().Info("CreateProduct", zap.Any("product", reqProduct))
	// zap.L().Debug("CreateProduct", zap.Any("product", product))


	product := FromReqCreateDTO(reqProduct)

	err := Repo().Create(product)	

	if err != nil {
		responseModel.ErrMsg = "Error while creating product"
		responseModel.ErrDsc = err.Error()
		responseModel.ResponseCode = http.StatusBadRequest
		c.JSON(responseModel.ResponseCode, responseModel)
		return
	}
	//this one is get Product by sku and return it with relations
	productDB,err := Repo().GetBySkuWithRelations(product.Sku)

	if err != nil {
		responseModel.ErrMsg = "Error while getting product"
		responseModel.ErrDsc = err.Error()
		responseModel.ResponseCode = http.StatusBadRequest
		c.JSON(responseModel.ResponseCode, responseModel)
		return
	}

	responseModel.ResponseCode = http.StatusOK
	responseModel.Data = productDB
	c.JSON(responseModel.ResponseCode, responseModel)

}
