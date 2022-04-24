package appAuth

import (

	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/config"
)

type authHandler struct {
	cfg *config.Config
}

func AuthHandler(r *gin.Engine, cfg *config.Config) {
	a := authHandler{
		cfg: cfg,
	}

	r.POST("/login", a.login)
	r.POST("/register", a.register)
	TokenControllerDef(r,cfg)

}



// func (a *authHandler) VerifyToken(c *gin.Context) {
// 	token := c.GetHeader("Authorization")
// 	decodedClaims := jwtHelper.VerifyToken(token, a.cfg.JWTConfig.SecretKey)

// 	c.JSON(http.StatusOK, decodedClaims)
// }
