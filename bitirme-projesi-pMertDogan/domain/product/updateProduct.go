package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain"
	"go.uber.org/zap"
)

//update product
func Update(c *gin.Context) {
	//Json
	c.Header("Content-Type", "application/json")
	var responseModel domain.ResponseModel

	//get  param
	id := c.Param("id")

	if id == "" {
		responseModel.ErrMsg = "id is empty"
		responseModel.ResponseCode = http.StatusBadRequest
		c.JSON(responseModel.ResponseCode, responseModel)
		zap.L().Debug("Update failed > id empty", zap.String("id", id))
		return
	}

	var req UpdateProductDTO



	err := c.ShouldBindJSON(&req)
	if err != nil {
		responseModel.ErrMsg = "unable bind body"
		responseModel.ErrDsc = err.Error()
		responseModel.ResponseCode = http.StatusInternalServerError
		c.JSON(responseModel.ResponseCode, responseModel)
		zap.L().Error("Update failed > unable bind body", zap.Error(err))
		return
	}
	//https://github.com/go-gorm/gorm/issues/4001 both struct should be same type
	x := Product{}
	copier.Copy(&x, &req)

	//PAtch
	v, err := Repo().Update(id,x)

	if err != nil {
		responseModel.ErrMsg = "patch failed"
		responseModel.ErrDsc = err.Error()
		responseModel.ResponseCode = http.StatusInternalServerError
		c.JSON(responseModel.ResponseCode, responseModel)
		zap.L().Error("Update failed  > patch failed", zap.Error(err))
		return
	}

	responseModel.Data = v
	responseModel.ResponseCode = http.StatusOK
	zap.L().Debug("Update success", zap.String("id", id))
	c.JSON(responseModel.ResponseCode, responseModel)

}
