package appAuth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/user"
)

func (a *authHandler) register(c *gin.Context) {

	//response is json
	c.Header("Content-Type", "application/json")
	var req user.RegisterRequestDTO
	var res user.ResponseModel
	//extract user from request with binding validation 
	if err := c.Bind(&req); err != nil {
		res.ErrMsg = "Your request body is not valid. Please check your request body."
		res.Err = err.Error()
		res.ResponseCode = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, res)
		return
	}

	//check is user exist with the same email adress
	 userR, err := user.Repo().CheckIsUserExistWithThisEmail(req.Email)
	//isUserExist return true if user exist on DB :)
	if err != nil {
		res.ErrMsg = "Something went wrong. Please try again later." 
		res.Err = err.Error()
		res.ResponseCode = http.StatusInternalServerError
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		// c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	if userR != nil {
		res.ErrMsg = "This email is already registered."
		res.ResponseCode = http.StatusBadRequest
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	//we can register user
	err = user.Repo().RegisterUser(req)

	if err != nil {
		res.ErrMsg = "Something went wrong. Please try again later."
		res.Err = err.Error()
		res.ResponseCode = http.StatusInternalServerError
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res.ResponseCode = http.StatusCreated
	res.UserData.Name = req.Name
	res.UserData.Email = req.Email


	c.JSON(http.StatusOK, res)

}
