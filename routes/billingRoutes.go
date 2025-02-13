package routes

import (
    "billing-microservice/controllers"
    "github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
    //Route to receive payment webhook
    r.POST("/webhook/payment", controllers.ReceivePayment)
}