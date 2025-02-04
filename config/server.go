package config

import (
    "github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitializeServer() {
    Router = gin.Default()
}
