package product

import (
	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/config"
	jwtUtils "github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/jwt"
)

func ProductControllerDef(router *gin.Engine, cfg *config.Config) {

	//https://github.com/gin-gonic/gin#using-middleware
	//Use JWT verification middleware

	product := router.Group("/product")
	//create MW 
	adminMW := jwtUtils.JWTAdminMiddleware(cfg.JWTConfig.SecretKey, cfg.JWTConfig.AccesTokenLifeMinute)


	product.GET("/", GetAllProductWithPagination)
	product.POST("/", adminMW, CreateProduct)
	product.POST("/bulk", adminMW, CreateBulkProduct)
	product.POST("/search", Search)
	product.DELETE("/:id", adminMW, Delete)
	product.PATCH("/:id", adminMW, Update)


}
