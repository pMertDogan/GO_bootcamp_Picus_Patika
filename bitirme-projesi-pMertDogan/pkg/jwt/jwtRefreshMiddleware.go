package jwtUtils

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"net/http"
)
/*
Handle request and check is they are admins and token is acces token
*/
func JWTRefreshMiddleware(secretKey string, accesTokenLifeMinute int64) gin.HandlerFunc {

	return func(c *gin.Context) {
		zap.L().Debug("JWTRefreshMiddleware is triggered")
		if c.GetHeader("Authorization") != "" {
			decodedClaims, err := VerifyDecodeToken(c.GetHeader("Authorization"), secretKey)

			//HANDLE JWT Errors
			if err != nil {
				c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
				c.Abort()
				return
			}

			if decodedClaims != nil {
				//ACCEPT Refresh TOKEN
				if decodedClaims.IsItAccesToken {
					c.JSON(http.StatusForbidden, gin.H{"error": "Please provide refresh token!"})
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
