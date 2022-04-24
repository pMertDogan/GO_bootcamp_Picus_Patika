package order

import (
	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/config"
	jwtUtils "github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/jwt"
)

func OrderControllerDef(router *gin.Engine, cfg *config.Config) {

	//https://github.com/gin-gonic/gin#using-middleware
	//Use JWT verification middleware

	order := router.Group("/order/:id")
	//Enable user ID check on order domain
	order.Use(jwtUtils.JWTUserIDMiddleware(cfg.JWTConfig.SecretKey, cfg.JWTConfig.AccesTokenLifeMinute))
    
	order.POST("/", CompleteOrder)
	order.GET("/", GetOrders)
	order.POST(":orderID", CancelOrder)
}
