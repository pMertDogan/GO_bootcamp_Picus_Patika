package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain"
	"go.uber.org/zap"
)

func Delete(c *gin.Context) {
	//Json
	c.Header("Content-Type", "application/json")

	//get  params
	id := c.Param("id")


	var responseModel domain.ResponseModel

	if id == "" {
		responseModel.ErrMsg = "id is empty"
		responseModel.ResponseCode = http.StatusBadRequest
		c.JSON(responseModel.ResponseCode, responseModel)
		zap.L().Debug("Delete > id empty", zap.String("id", id))
		return
	}

	v,err:= Repo().Delete(id)

	if err != nil {
		responseModel.ErrMsg = "error getting products"
		responseModel.ErrDsc = err.Error()
		responseModel.ResponseCode = http.StatusInternalServerError
		c.JSON(responseModel.ResponseCode, responseModel)
		zap.L().Error("Delete > error getting products", zap.Error(err))
		return
	}

	responseModel.Data = v
	responseModel.ResponseCode = http.StatusOK
	c.JSON(responseModel.ResponseCode, responseModel)

}
