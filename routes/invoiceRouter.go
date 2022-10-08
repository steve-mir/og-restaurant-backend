package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/steve-mir/og-restaurant-backend/controllers"
)

func InvoiceRoutes(route *gin.Engine) {
	route.GET("/invoices", controllers.InvoiceList())
	route.GET("/invoices/:id", controllers.InvoiceGet())
	route.POST("/invoices", controllers.InvoiceCreate())
	route.PATCH("/invoices/:id", controllers.InvoiceUpdate())
	route.DELETE("/invoices/:id", controllers.InvoiceDelete())
}
