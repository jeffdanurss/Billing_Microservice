package routes

import (
	"billing-microservice/controllers"
	"billing-microservice/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/billing", middlewares.AuthMiddleware(), controllers.CrearFactura)
	r.GET("/billing/:id/pdf", middlewares.AuthMiddleware(), controllers.DescargarFactura)
}
