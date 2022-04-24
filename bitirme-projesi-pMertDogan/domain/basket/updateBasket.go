package basket

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/product"
)

func UpdateBasket(c *gin.Context) {

	response := domain.ResponseModel{}

	//get userID from url
	userID := c.Param("id")

	//checked on MW side but we can add addtionale check here
	if userID == "" {
		response.ErrMsg = "userID is required"
		response.ResponseCode = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//userID to int
	userIDInt, err := strconv.Atoi(userID)

	if err != nil {
		response.ErrMsg = "userID must be integer"
		response.ResponseCode = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, response)
		return

	}

	var req UpdateBasketDTO
	//bind to dto
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ResponseCode = http.StatusBadRequest
		response.ErrMsg = "error binding json "
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//Check if basketID is valid
	oldBasket, err := Repo().GetBasketByUserIDAndID(userIDInt, req.BasketID)

	if err != nil {
		response.ErrMsg = "Basket not found"
		response.ResponseCode = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//check total quantity is not less than 1
	if req.TotalQuantity < 1 {
		response.ErrMsg = "Total quantity must be greater than 0"
		response.ResponseCode = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, response)
		return
	}

	productQuantity, err := product.Repo().GetProductQuantityById(oldBasket.ProductID)

	if err != nil {
		response.ResponseCode = http.StatusBadRequest
		response.ErrMsg = "product not found"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//check if product has enought stock
	if productQuantity < req.TotalQuantity {
		response.ResponseCode = http.StatusBadRequest
		response.ErrMsg = "not enought stock"
		response.ErrDsc = "product quantity is " + strconv.Itoa(productQuantity)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//update basket
	err = Repo().UpdateBasketQuantity(userIDInt, req.TotalQuantity, req.BasketID)

	if err != nil {
		response.ErrMsg = "Error updating basket"
		response.ResponseCode = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, response)
		return
	}

	baskets, err := Repo().GetBasketsByUserIDWithPaginations(userIDInt,"0","50")

	//return success
	response.ResponseCode = http.StatusOK
	response.Data = baskets
	response.PageNo = "0"
	response.PageSize = "50"
	c.JSON(http.StatusOK, response)

}
