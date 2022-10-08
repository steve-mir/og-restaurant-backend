package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/steve-mir/og-restaurant-backend/controllers"
)

func TableRoutes(route *gin.Engine) {
	route.GET("/tables", controllers.GetTables())
	route.GET("/tables/:id", controllers.GetTable())
	route.POST("/tables", controllers.CreateTable())
	route.PATCH("/table/:id", controllers.UpdateTable())
}
