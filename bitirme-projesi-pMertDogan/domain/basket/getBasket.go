package basket

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain"
	"go.uber.org/zap"
)

func GetBasket(c *gin.Context) {

	response := domain.ResponseModel{}


	//get userID from url
	userID := c.Param("id")

	//get page from url
	page := c.DefaultQuery("page","0")
	//get pageSize from url
	pageSize := c.DefaultQuery("pageSize","50")

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

	//add to Basket
	v,err := Repo().GetBasketsByUserIDWithPaginations(userIDInt,page,pageSize)

	if err != nil {
		response.ErrMsg = "error getting basket"
		response.ResponseCode = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, response)
		return
	}
	zap.L().Debug("basket getted")

	//return success
	response.ResponseCode = http.StatusOK
	response.Data = v
	response.PageNo = page
	response.PageSize = pageSize
	c.JSON(http.StatusOK, response)

}
