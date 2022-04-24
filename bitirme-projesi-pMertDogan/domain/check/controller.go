package check

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/database"
	"go.uber.org/zap"
)

func CheckControllerDef(router *gin.Engine) {

	c := router.Group("/check")
	{

		c.GET("/status", statusCheck)
		c.GET("/ready", ready)
		// c.POST("/read", readEndpoint)
	}
}

func statusCheck(c *gin.Context) {
	zap.L().Info("status check called")
	c.JSON(http.StatusOK, gin.H{
		"message": "service is runnng",
	})
}

func ready(c *gin.Context) {

	zap.L().Info("ready is called")
	openCon, err := database.StatusCheck()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "There is a problem with DB connection",
		})
	}
	

	c.JSON(http.StatusOK, gin.H{
		"message": "service is runnng",
		//Its not good idea to provide database connection count to client
		"DBOpenConnectionCount ": openCon,
	})
}
