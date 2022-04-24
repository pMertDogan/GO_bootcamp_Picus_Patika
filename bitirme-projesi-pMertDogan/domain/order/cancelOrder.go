package order

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain"
	"go.uber.org/zap"
)

//Get customer orders with pagination
func CancelOrder(c *gin.Context) {
	//init response model
	response := domain.ResponseModel{}

	//get userID from Param
	id := c.Param("id")
	orderID := c.Param("orderID")
	zap.L().Debug("ID is", zap.String("id", id))
	zap.L().Debug("orderID is", zap.String("oderID", orderID))

	//Check is provided ids are int
	//convert id to int
	userID, err := strconv.Atoi(id)
	if err != nil {
		//verify sended one is int
		response.ErrMsg = "Cannot convert id to int"
		response.ErrDsc = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//convert orderID to int
	orderIDInt, err := strconv.Atoi(orderID)
	if err != nil {
		//verify sended one is int
		response.ErrMsg = "Cannot convert orderid to int"
		response.ErrDsc = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//check If user has permission to cancel order
	//get orders from userID
	order, err := Repo().HasOrder(userID, orderIDInt)

	if err != nil {

		response.ErrMsg = "User does not have permission to cancel order or order does not exist for this user"
		response.ErrDsc = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//check order createdAT time is not older than 14 days
	if order.CreatedAt.AddDate(0, 0, 14).Before(time.Now()) {
		//cancel Time Period is over
		response.ErrMsg = "Cancel  Period is over, user cant cancel order"
		response.ErrDsc = "This order is created at " + order.CreatedAt.String() + ". Can't be canceled after 14 days . Total days after order creation is " + strconv.Itoa(int(time.Now().Sub(order.CreatedAt).Hours()/24)) + " and current date is " + time.Now().String()
		c.JSON(http.StatusNotAcceptable, response)
		return
	}

	if err != nil {
		response.ErrMsg = "Cannot find order"
		response.ErrDsc = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//cancel order by orderID verify with userID
	err = Repo().CancelOrder(*order)

	//return succes as response
	response.Data = order
	response.ResponseCode = http.StatusOK
	c.JSON(http.StatusOK, response)
}
