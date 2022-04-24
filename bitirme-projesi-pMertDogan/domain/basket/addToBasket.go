package basket

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/product"
)

// AddToBasket adds product to basket
func AddToBasket(c *gin.Context) {

	response := domain.ResponseModel{}

	//get userID from url
	userID := c.Param("id")

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

	var req AddToBasketDTO
	//bind to dto
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ResponseCode = http.StatusBadRequest
		response.ErrMsg = "error binding json "
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//check if product has enought stock
	productQuantity, err := product.Repo().GetProductQuantityById(req.ProductID)

	if err != nil {
		response.ResponseCode = http.StatusBadRequest
		response.ErrMsg = "product not found"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//get basket by quantity if exist
	basket, _ := Repo().GetBasketByUserIDAndProductID(userIDInt, req.ProductID)

	total := req.TotalQuantity
	if basket != nil {
		total += basket.TotalQuantity
	}

	//or we can we can define max allowed quantity = 10
	// if productQuantity < 10+basket.TotalQuantity {
	if productQuantity < total {
		response.ResponseCode = http.StatusBadRequest
		response.ErrMsg = "not enought stock"
		response.ErrDsc = "product quantity is " + strconv.Itoa(productQuantity)
		c.JSON(http.StatusBadRequest, response)
		return
	}


	//add to Basket or create new one
	err = Repo().CreateOrUpdateBasket(userIDInt, req.ProductID, req.TotalQuantity)

	if err != nil {
		response.ResponseCode = http.StatusBadRequest
		response.ErrMsg = "Unable add to basket"
		response.ErrDsc = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}


	v , err := Repo().GetBasketsByUserIDWithPaginations(userIDInt,"0","50")
	//return success
	response.ResponseCode = http.StatusOK
	response.Data = v
	response.PageNo = "0"
	response.PageSize = "50"
	c.JSON(http.StatusOK, response)



}
