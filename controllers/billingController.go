package controllers

import (
	"billing-microservice/config"
	"billing-microservice/models"
	"billing-microservice/services"
	"billing-microservice/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create a new billing
func CrearFactura(c *gin.Context) {
	var factura models.Factura
	if err := c.ShouldBindJSON(&factura); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// save billing into dataset
	result := config.DB.Create(&factura)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	err := services.SendWebhookNotification(factura)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al enviar notificación"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Factura creada exitosamente", "factura": factura})
}

func DescargarFactura(c *gin.Context) {
	facturaID := c.Param("id")
	var factura models.Factura


	result := config.DB.First(&factura, facturaID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Factura no encontrada"})
		return
	}


	filename := fmt.Sprintf("factura_%d.pdf", factura.OrdenID)
	err := utils.GenerateFacturaPDF(factura, filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar PDF"})
		return
	}


	c.File(filename)
}
