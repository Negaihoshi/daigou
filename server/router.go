package server

import (
	"github.com/gin-gonic/gin"
	"github.com/negaihoshi/daigou/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)
	// router.Use(middlewares.AuthMiddleware())

	api := router.Group("api")
	{
		orderGroup := api.Group("order")
		{
			order := new(controllers.OrderController)
			orderGroup.GET("", order.Index)
			orderGroup.POST("", order.Store)
			orderGroup.PUT("", order.Update)
			orderGroup.DELETE("", order.Destroy)
		}
	}
	return router

}
