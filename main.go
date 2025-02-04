package main

import (
	"billing-microservice/config"
	"billing-microservice/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Conect to database
	config.Connect()

	// Create a new GinCreate route
	r := gin.Default()

	// setting routes
	routes.RegisterRoutes(r)

	// set up server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor: ", err)
	}
}
