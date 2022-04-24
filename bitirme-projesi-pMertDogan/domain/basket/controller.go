package basket

import (
	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/config"
	jwtUtils "github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/jwt"
)

func BasketControllerDef(router *gin.Engine, cfg *config.Config) {

	//https://github.com/gin-gonic/gin#using-middleware
	//Use JWT verification middleware

	userIDMW := jwtUtils.JWTUserIDMiddleware(cfg.JWTConfig.SecretKey, cfg.JWTConfig.AccesTokenLifeMinute)
	// adminMW := jwtUtils.JWTCheckAcessTokenMiddleware(cfg.JWTConfig.SecretKey, cfg.JWTConfig.AccesTokenLifeMinute)
	Basket := router.Group("/basket")
	//user should be provide acces token for this route

	Basket.POST("/:id", userIDMW, AddToBasket)
	Basket.GET("/:id", userIDMW, GetBasket)
	Basket.PATCH("/:id", userIDMW, UpdateBasket)
	Basket.DELETE("/:id/:basketID", userIDMW, DeleteBasket)
	// cat.GET("/" ,GetAllCategoriesWithPagination)
}
