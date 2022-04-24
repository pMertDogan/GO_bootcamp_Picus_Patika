package jwtUtils

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"net/http"
)
/*
Handle request and check is acces token is valid
*/
func JWTCheckAcessTokenMiddleware(secretKey string, accesTokenLifeMinute int64) gin.HandlerFunc {

	return func(c *gin.Context) {
		zap.L().Debug("JWTAcessTokenMiddleware is triggered")
		//Get token from header
		if c.GetHeader("Authorization") != "" {
			decodedClaims, err := VerifyDecodeToken(c.GetHeader("Authorization"), secretKey)
			//HANDLE JWT Errors
			if err != nil {
				c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
				c.Abort()
				return
			}

			if decodedClaims != nil {
				//ACCEPT acces TOKEN
				if !decodedClaims.IsItAccesToken {
					c.JSON(http.StatusForbidden, gin.H{"error": "Please provide acces token!"})
					c.Abort()
					return
				}
				c.Next()
				return
			}
			c.JSON(http.StatusForbidden, gin.H{"error": "Token is invalid! Unable parse token!"})
			c.Abort()
			return

		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Please provide token!"})
		}
		c.Abort()
		return
	}
}
