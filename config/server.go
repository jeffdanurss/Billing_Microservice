package config

import (
    "github.com/gin-gonic/gin"
    "billing-microservice/routes"
)

var Router *gin.Engine

func InitializeServer() {
    Router = gin.Default()
    routes.RegisterRoutes(Router) // Registra las rutas
}