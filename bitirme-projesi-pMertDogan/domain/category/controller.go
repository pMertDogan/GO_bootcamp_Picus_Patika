package category

import (
	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/config"
	jwtUtils "github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/jwt"
)

func CategoryControllerDef(router *gin.Engine, cfg *config.Config) {

	//https://github.com/gin-gonic/gin#using-middleware
	//Use JWT verification middleware

	cat := router.Group("/category")

	cat.POST("/upload", jwtUtils.JWTAdminMiddleware(cfg.JWTConfig.SecretKey, cfg.JWTConfig.AccesTokenLifeMinute), UploadCategorysFromCSV)
	cat.GET("/" ,GetAllCategoriesWithPagination)
}
