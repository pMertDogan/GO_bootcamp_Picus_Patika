package category

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain"
)

func GetAllCategoriesWithPagination(c *gin.Context) {

	response := domain.ResponseModel{}
	//get query params
	pageSize := c.DefaultQuery("pageSize", "10")
	pageNo := c.DefaultQuery("pageNo", "1")

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

	

	//get all categories with pagination
	v, err := Repo().GetAllCategoriesWithPagination(pageNoInt, pageSizeInt)

	if err != nil {
		response.ResponseCode = http.StatusInternalServerError
		response.ErrMsg = "error getting  categories with pagination"
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	//No data found
	if len(v) == 0 {
		response.ResponseCode = http.StatusNotFound
		response.ErrMsg = "no data found"
		c.JSON(http.StatusNotFound, response)
		return
	}

	response.ResponseCode = http.StatusOK
	response.Data = v
	//To fix the problem if user send pageSize > 100. We will return only 100
	response.PageNo ,response.PageSize = domain.CalcPageAndSizeReturnString(pageNoInt, pageSizeInt)
	c.JSON(http.StatusOK, response)

}


