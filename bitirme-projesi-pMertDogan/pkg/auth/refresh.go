package appAuth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/config"
	jwtUtils "github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/jwt"
)

type res struct {
	AccesToken   string `json:"accesToken,omitempty"`
	ResponseCode int    `json:"responseCode"`
	Error        error  `json:"error,omitempty"`
}

func TokenControllerDef(router *gin.Engine, cfg *config.Config) {

	var res res

	//https://github.com/gin-gonic/gin#using-middleware
	//Use JWT verification middleware

	tok := router.Group("/token")

	//check if  refresh token is valid
	tok.POST("/fresh", jwtUtils.JWTRefreshMiddleware(cfg.JWTConfig.SecretKey, cfg.JWTConfig.AccesTokenLifeMinute), func(c *gin.Context) {
		
		// if c.GetHeader("Authorization") != "" {
			user, _ := jwtUtils.VerifyDecodeToken(c.GetHeader("Authorization"), cfg.JWTConfig.SecretKey)

			
			jwtClaimsAccess := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"userId":         user.UserId,
				"email":          user.Email,
				"iat":            time.Now().Unix(),                                                                      //issued at current time
				"exp":            time.Now().Add(time.Duration(cfg.JWTConfig.AccesTokenLifeMinute) * time.Minute).Unix(), //expiration time is one hour
				"isAdmin":        user.IsAdmin,
				"isItAccesToken": true,
			})
			// //create JWT token

			accesToken := jwtUtils.GenerateToken(jwtClaimsAccess, cfg.JWTConfig.SecretKey)

			res.AccesToken = accesToken

			res.ResponseCode = http.StatusOK
			c.JSON(http.StatusOK, res)
		// }

	})
}
