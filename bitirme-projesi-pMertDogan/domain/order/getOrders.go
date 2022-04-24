package order

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain"
	"go.uber.org/zap"
)

//Get customer orders with pagination
func GetOrders(c *gin.Context) {
	//init response model
	response := domain.ResponseModel{}

	//get userID from Param
	id := c.Param("id")
	pageSize := c.DefaultQuery("pageSize", "10")
	pageNo := c.DefaultQuery("pageNo", "1")
	zap.L().Debug("ID is", zap.String("id", id))

	//convert pageSize to int
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		//verify sended one is int
		response.ErrMsg = "Cannot convert pageSize to int"
		response.ErrDsc = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//convert pageNo to int
	pageNoInt, err := strconv.Atoi(pageNo)
	if err != nil {
		//verify sended one is int
		response.ErrMsg = "Cannot convert pageNo to int"
		response.ErrDsc = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}


	//convert id to int
	userID, err := strconv.Atoi(id)
	if err != nil {
		response.ErrMsg = "Cannot convert id to int"
		response.ErrDsc = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//get orders from userID
	orders, err := Repo().GetOrders(userID,pageNoInt,pageSizeInt)

	
	//return succes as response
	response.Data = orders
	response.ResponseCode = http.StatusOK
	response.PageNo = pageNo
	response.PageSize = pageSize
	c.JSON(http.StatusOK, response)
}
