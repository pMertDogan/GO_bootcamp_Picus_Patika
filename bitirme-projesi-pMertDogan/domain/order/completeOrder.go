package order

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/basket"
	"go.uber.org/zap"
)

func CompleteOrder(c *gin.Context) {
	//init response model
	response := domain.ResponseModel{}

	//get userID from Param
	id := c.Param("id")

	zap.L().Debug("ID is", zap.String("id", id))

	//convert id to int
	userID, err := strconv.Atoi(id)
	if err != nil {
		response.ErrMsg = "Cannot convert id to int"
		response.ErrDsc = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//init request model
	var req CompleteOrderDTO
	//bind to request model
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrMsg = "Invalid request body"
		response.ErrDsc = err.Error()
		response.ResponseCode = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//check is basketsid are correct and valid for userID
	err = req.ValideteBasketIDs(userID)

	if err != nil {
		response.ErrMsg = "Invalid basketIDs"
		response.ErrDsc = err.Error()
		response.ResponseCode = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//get baskets from basketIDs
	baskets, err := basket.Repo().GetBasketsByUserIDAndBasketIDs(req.BasketIDs)
	zap.L().Info("baskets", zap.Any("baskets are ", baskets))

	if err != nil {
		response.ErrMsg = "Cannot get baskets"
		response.ErrDsc = err.Error()
		response.ResponseCode = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//complete order
	err = Repo().CompleteOrder(baskets, req.Comment, req.ShipingAddress, req.BillinAddress)

	//At this point we can say user send correct request with correct data and we can continue
	//start #transaction and #lock specific row inside product table
	//Lets check is current quantity of each product is enough for order

	if err != nil {
		response.ErrMsg = "Cannot complete order"
		response.ErrDsc = err.Error()
		response.ResponseCode = http.StatusNotAcceptable
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//return succes as response
	response.ResponseCode = http.StatusOK
	c.JSON(http.StatusOK, response)
}
