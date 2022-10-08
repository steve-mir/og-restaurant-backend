package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/steve-mir/og-restaurant-backend/controllers"
)

func MenuRoutes(route *gin.Engine) {
	route.GET("/menu", controllers.GetMenus())
	route.GET("/menu/:id", controllers.GetMenu())
	route.POST("/menu", controllers.CreateMenu())
	route.PUT("/menu/:id", controllers.UpdateMenu())
	route.DELETE("/menu/", controllers.DeleteMenu())
}
