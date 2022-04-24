package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UploadCategorysFromCSV(c *gin.Context) {
	// router.POST("/upload", func(c *gin.Context) {

		response := ResponseModel{}

		// single-file upload read categoryFile
		file, _ := c.FormFile("categoryFile")

		//IF user not provide file
		if file == nil {
			zap.L().Error("categoryFile is nil")
			response.ResponseCode = http.StatusBadRequest
			response.ErrMsg = "categoryFile is nil"
			c.JSON(http.StatusBadRequest, response)
			return
		}

		zap.L().Info("Uploaded succes", zap.String("fileName ", file.Filename))

		//send file to our converter function
		// https://stackoverflow.com/questions/40956103/how-to-convert-multipart-fileheader-file-type-to-os-file-in-golang
		//TLDR; we dont need :)
		csvCategories, err := CategoryFromCSV(file)
		if err != nil {
			response.ErrMsg = "check the csv  file. " + err.Error()
			response.ResponseCode = http.StatusBadRequest
			zap.L().Error("check the csv  file. " + err.Error())
			c.JSON(http.StatusBadRequest, response)
			return
		}

		//call category repo and save it to db
		Repo().CreateCategories(csvCategories)

		//display 50x categories
		v, err := Repo().GetAllCategoriesWithLimit(50)
		if err != nil {
			zap.L().Error("error getting all categories", zap.Error(err))
			response.ResponseCode = http.StatusInternalServerError
			response.ErrMsg = "error getting all categories"
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		
		response.Data = v
		c.JSON(http.StatusOK, response)
}