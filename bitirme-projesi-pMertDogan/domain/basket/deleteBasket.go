package basket

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain"
)

func DeleteBasket(c *gin.Context) {

	response := domain.ResponseModel{}

	//get userID from url
	userID := c.Param("id")
	basketID := c.Param("basketID")

	//checked on MW side but we can add addtionale check here
	if userID == "" || basketID == "" {
		response.ErrMsg = "userID  and basketID is required"
		response.ResponseCode = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//userID to int
	userIDInt, err := strconv.Atoi(userID)
	basketIDInt , err := strconv.Atoi(basketID)

	if err != nil {
		response.ErrMsg = "userID  and basketID must be integer"
		response.ResponseCode = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, response)
		return

	}

	_,err = Repo().GetBasketByUserIDAndID(userIDInt,basketIDInt)
	
	if err != nil {
		response.ErrMsg = "Basket not found"
		response.ResponseCode = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, response)
		return
	}

	

	//update basket
	err = Repo().RemoveBasketByUserIDBasketID(userIDInt,basketIDInt)

	if err != nil {
		response.ErrMsg = "Error deleting basket"
		response.ResponseCode = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, response)
		return
	}

	baskets ,err := Repo().GetBasketsByUserIDWithPaginations(userIDInt,"0","50")

	//return success
	response.ResponseCode = http.StatusOK
	response.Data = baskets
	response.PageNo = "0"
	response.PageSize = "50"
	c.JSON(http.StatusOK, response)

}
