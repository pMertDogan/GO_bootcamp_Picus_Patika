package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain"
	"go.uber.org/zap"
)

func CreateBulkProduct(c *gin.Context) {

		response := domain.ResponseModel{}

		// single-file upload read productsFile
		file, _ := c.FormFile("products")

		//IF user not provide file
		if file == nil {
			zap.L().Error("products File is nil")
			response.ResponseCode = http.StatusBadRequest
			response.ErrMsg = "products File is nil"
			c.JSON(http.StatusBadRequest, response)
			return
		}

		zap.L().Info("Uploaded succes", zap.String("fileName ", file.Filename))

		//send file to our converter function
		// https://stackoverflow.com/questions/40956103/how-to-convert-multipart-fileheader-file-type-to-os-file-in-golang
		//TLDR; we dont need :)
		csvProducts, err := ProductFromCSV(file)
		if err != nil {
			response.ErrMsg = "check the csv  file. " + err.Error()
			response.ResponseCode = http.StatusBadRequest
			zap.L().Error("check the csv  file. " + err.Error())
			c.JSON(http.StatusBadRequest, response)
			return
		}

		//call products repo and save it to db
		
		Repo().CreateBulkProduct(csvProducts)

		v, err := Repo().GetAllWithPagination(1, 100)
		if err != nil {
			zap.L().Error("error getting all categories", zap.Error(err))
			response.ResponseCode = http.StatusInternalServerError
			response.ErrMsg = "error getting all categories"
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		response.PageSize = "100"
		response.Data = v
		response.ResponseCode = http.StatusOK
		
		c.JSON(http.StatusOK, response)
}
