package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/steve-mir/og-restaurant-backend/controllers"
)

func FoodRoutes(route *gin.Engine) {
	route.GET("/foods", controllers.GetFoods())
	route.GET("/foods/:id", controllers.GetFood())
	route.POST("/foods", controllers.CreateFood())
	route.PATCH("foods/:id", controllers.UpdateFood())
}
