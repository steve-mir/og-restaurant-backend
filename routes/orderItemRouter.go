package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/steve-mir/og-restaurant-backend/controllers"
)

func OrderItemRoutes(route *gin.Engine) {
	route.GET("/orderItems", controllers.GetOrderItems())
	route.GET("/orderItems/:id", controllers.GetOrderItem())
	route.GET("/orderItems-order/:id", controllers.GetOrderItemByOrder())
	route.POST("/orderItems", controllers.CreateOrderItem())
	route.PUT("/orderItems/:id", controllers.UpdateOrderItem())
	route.DELETE("/orderItems/:id", controllers.DeleteOrderItem())
}
