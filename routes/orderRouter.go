package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/steve-mir/og-restaurant-backend/controllers"
)

func OrderRoutes(route *gin.Engine) {
	route.GET("/orders", controllers.GetOrders())
	route.GET("/orders/:id", controllers.GetOrder())
	route.POST("/orders", controllers.CreateOrder())
	route.PUT("/orders/:id", controllers.UpdateOrder())
	route.DELETE("/orders/:id", controllers.DeleteOrder())
}
