package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/steve-mir/og-restaurant-backend/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controllers.GetUser())
	incomingRoutes.GET("/users/:id", controllers.GetUser())
	incomingRoutes.POST("/users/register", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
}
